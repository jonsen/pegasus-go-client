// Copyright (c) 2017, Xiaomi, Inc.  All rights reserved.
// This source code is licensed under the Apache License Version 2.0, which
// can be found in the LICENSE file in the root directory of this source tree.

package rpc

import (
	"net"
	"testing"
	"time"

	"context"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
)

// This test ensures that:
// - RpcConn can be reopened after it closed.
func TestRpcConn_CreateConnected(t *testing.T) {
	defer leaktest.Check(t)()

	for i := 0; i < 3; i++ {
		conn := NewRpcConn("www.baidu.com:80")
		err := conn.TryConnect()

		assert.Nil(t, err)
		assert.Equal(t, conn.cstate, ConnStateReady)
		assert.NotNil(t, conn.rstream)
		assert.NotNil(t, conn.wstream)

		conn.Close()
	}
}

// Ensure that read write from a not-ready connection returns ErrConnectionNotReady.
func TestRpcConn_ReadWriteNotReady(t *testing.T) {
	defer leaktest.Check(t)()

	conn := NewRpcConn("www.baidu.com:80")
	err := conn.TryConnect()
	assert.Nil(t, err)

	conn.Close()

	go func() {
		_, err := conn.Read(4)
		assert.Equal(t, err, ErrConnectionNotReady)
	}()

	go func() {
		err := conn.Write([]byte("ping"))
		assert.Equal(t, err, ErrConnectionNotReady)
	}()
}

// Ensure that a blocked read can be cancelled by closing the RpcConn.
func TestRpcConn_ReadCancelled(t *testing.T) {
	defer leaktest.Check(t)()

	conn := NewRpcConn("www.baidu.com:80")
	err := conn.TryConnect()
	assert.Nil(t, err)

	go func() {
		// we can never read from baidu if we didn't write anything.
		conn.Read(4)
	}()

	time.Sleep(time.Second)
	conn.Close()
}

func TestRpcConn_NewRpcConnectFailed(t *testing.T) {
	defer leaktest.CheckTimeout(t, time.Second*6)()

	// it must time out.
	conn := NewRpcConn("www.baidu.com:12321")
	err := conn.TryConnect()
	assert.NotNil(t, err)
	assert.Equal(t, ConnStateTransientFailure, conn.cstate)
	conn.Close()
}

// This test verifies that a connecting RpcConn can be cancelled
// immediately by Close.
func TestRpcConn_CancelConnecting(t *testing.T) {
	defer leaktest.Check(t)()

	conn := NewRpcConn("www.baidu.com:12321")
	go func() {
		conn.TryConnect()
	}()

	time.Sleep(time.Second)
	assert.Equal(t, ConnStateConnecting, conn.cstate)
	conn.Close()
}

func TestRpcConn_WriteAndRead(t *testing.T) {
	defer leaktest.Check(t)()

	// start echo server first
	conn := NewRpcConn("0.0.0.0:8800")
	defer conn.Close()

	assert.Nil(t, conn.TryConnect())

	data := []byte("ping")
	assert.Nil(t, conn.Write(data))

	actual, err := conn.Read(4)
	assert.Nil(t, err)
	assert.Equal(t, data, actual)

	// Ensure that read will timeout after 1 second.
	_, err = conn.Read(1)
	opErr, ok := err.(*net.OpError)
	assert.True(t, ok)
	assert.True(t, opErr.Timeout())

	// Ensure read can restart from last failure.
	data = []byte("hello")
	assert.Nil(t, conn.Write(data))
	actual, err = conn.Read(5)
	assert.Nil(t, err)
	assert.Equal(t, data, actual)
}

func Test_IsNetworkTimeoutErr(t *testing.T) {
	// timeout error but not a network error
	assert.False(t, IsNetworkTimeoutErr(context.DeadlineExceeded))

	err := NewRpcConn("www.baidu.com:12321").TryConnect()
	assert.True(t, IsNetworkTimeoutErr(err))
}
