// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/pegasus-kv/pegasus-go-client/idl/rrdb"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  update_response put(update_request update)")
	fmt.Fprintln(os.Stderr, "  update_response multi_put(multi_put_request request)")
	fmt.Fprintln(os.Stderr, "  update_response remove(blob key)")
	fmt.Fprintln(os.Stderr, "  multi_remove_response multi_remove(multi_remove_request request)")
	fmt.Fprintln(os.Stderr, "  read_response get(blob key)")
	fmt.Fprintln(os.Stderr, "  multi_get_response multi_get(multi_get_request request)")
	fmt.Fprintln(os.Stderr, "  count_response sortkey_count(blob hash_key)")
	fmt.Fprintln(os.Stderr, "  ttl_response ttl(blob key)")
	fmt.Fprintln(os.Stderr, "  scan_response get_scanner(get_scanner_request request)")
	fmt.Fprintln(os.Stderr, "  scan_response scan(scan_request request)")
	fmt.Fprintln(os.Stderr, "  void clear_scanner(i64 context_id)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := rrdb.NewRrdbClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "put":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Put requires 1 args")
			flag.Usage()
		}
		arg27 := flag.Arg(1)
		mbTrans28 := thrift.NewTMemoryBufferLen(len(arg27))
		defer mbTrans28.Close()
		_, err29 := mbTrans28.WriteString(arg27)
		if err29 != nil {
			Usage()
			return
		}
		factory30 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt31 := factory30.GetProtocol(mbTrans28)
		argvalue0 := rrdb.NewUpdateRequest()
		err32 := argvalue0.Read(jsProt31)
		if err32 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Put(value0))
		fmt.Print("\n")
		break
	case "multi_put":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MultiPut requires 1 args")
			flag.Usage()
		}
		arg33 := flag.Arg(1)
		mbTrans34 := thrift.NewTMemoryBufferLen(len(arg33))
		defer mbTrans34.Close()
		_, err35 := mbTrans34.WriteString(arg33)
		if err35 != nil {
			Usage()
			return
		}
		factory36 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt37 := factory36.GetProtocol(mbTrans34)
		argvalue0 := rrdb.NewMultiPutRequest()
		err38 := argvalue0.Read(jsProt37)
		if err38 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MultiPut(value0))
		fmt.Print("\n")
		break
	case "remove":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Remove requires 1 args")
			flag.Usage()
		}
		arg39 := flag.Arg(1)
		mbTrans40 := thrift.NewTMemoryBufferLen(len(arg39))
		defer mbTrans40.Close()
		_, err41 := mbTrans40.WriteString(arg39)
		if err41 != nil {
			Usage()
			return
		}
		factory42 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt43 := factory42.GetProtocol(mbTrans40)
		argvalue0 := rrdb.NewBlob()
		err44 := argvalue0.Read(jsProt43)
		if err44 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Remove(value0))
		fmt.Print("\n")
		break
	case "multi_remove":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MultiRemove requires 1 args")
			flag.Usage()
		}
		arg45 := flag.Arg(1)
		mbTrans46 := thrift.NewTMemoryBufferLen(len(arg45))
		defer mbTrans46.Close()
		_, err47 := mbTrans46.WriteString(arg45)
		if err47 != nil {
			Usage()
			return
		}
		factory48 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt49 := factory48.GetProtocol(mbTrans46)
		argvalue0 := rrdb.NewMultiRemoveRequest()
		err50 := argvalue0.Read(jsProt49)
		if err50 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MultiRemove(value0))
		fmt.Print("\n")
		break
	case "get":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Get requires 1 args")
			flag.Usage()
		}
		arg51 := flag.Arg(1)
		mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
		defer mbTrans52.Close()
		_, err53 := mbTrans52.WriteString(arg51)
		if err53 != nil {
			Usage()
			return
		}
		factory54 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt55 := factory54.GetProtocol(mbTrans52)
		argvalue0 := rrdb.NewBlob()
		err56 := argvalue0.Read(jsProt55)
		if err56 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Get(value0))
		fmt.Print("\n")
		break
	case "multi_get":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MultiGet requires 1 args")
			flag.Usage()
		}
		arg57 := flag.Arg(1)
		mbTrans58 := thrift.NewTMemoryBufferLen(len(arg57))
		defer mbTrans58.Close()
		_, err59 := mbTrans58.WriteString(arg57)
		if err59 != nil {
			Usage()
			return
		}
		factory60 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt61 := factory60.GetProtocol(mbTrans58)
		argvalue0 := rrdb.NewMultiGetRequest()
		err62 := argvalue0.Read(jsProt61)
		if err62 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MultiGet(value0))
		fmt.Print("\n")
		break
	case "sortkey_count":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SortkeyCount requires 1 args")
			flag.Usage()
		}
		arg63 := flag.Arg(1)
		mbTrans64 := thrift.NewTMemoryBufferLen(len(arg63))
		defer mbTrans64.Close()
		_, err65 := mbTrans64.WriteString(arg63)
		if err65 != nil {
			Usage()
			return
		}
		factory66 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt67 := factory66.GetProtocol(mbTrans64)
		argvalue0 := rrdb.NewBlob()
		err68 := argvalue0.Read(jsProt67)
		if err68 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.SortkeyCount(value0))
		fmt.Print("\n")
		break
	case "ttl":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TTL requires 1 args")
			flag.Usage()
		}
		arg69 := flag.Arg(1)
		mbTrans70 := thrift.NewTMemoryBufferLen(len(arg69))
		defer mbTrans70.Close()
		_, err71 := mbTrans70.WriteString(arg69)
		if err71 != nil {
			Usage()
			return
		}
		factory72 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt73 := factory72.GetProtocol(mbTrans70)
		argvalue0 := rrdb.NewBlob()
		err74 := argvalue0.Read(jsProt73)
		if err74 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TTL(value0))
		fmt.Print("\n")
		break
	case "get_scanner":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetScanner requires 1 args")
			flag.Usage()
		}
		arg75 := flag.Arg(1)
		mbTrans76 := thrift.NewTMemoryBufferLen(len(arg75))
		defer mbTrans76.Close()
		_, err77 := mbTrans76.WriteString(arg75)
		if err77 != nil {
			Usage()
			return
		}
		factory78 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt79 := factory78.GetProtocol(mbTrans76)
		argvalue0 := rrdb.NewGetScannerRequest()
		err80 := argvalue0.Read(jsProt79)
		if err80 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetScanner(value0))
		fmt.Print("\n")
		break
	case "scan":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Scan requires 1 args")
			flag.Usage()
		}
		arg81 := flag.Arg(1)
		mbTrans82 := thrift.NewTMemoryBufferLen(len(arg81))
		defer mbTrans82.Close()
		_, err83 := mbTrans82.WriteString(arg81)
		if err83 != nil {
			Usage()
			return
		}
		factory84 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt85 := factory84.GetProtocol(mbTrans82)
		argvalue0 := rrdb.NewScanRequest()
		err86 := argvalue0.Read(jsProt85)
		if err86 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Scan(value0))
		fmt.Print("\n")
		break
	case "clear_scanner":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ClearScanner requires 1 args")
			flag.Usage()
		}
		argvalue0, err87 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err87 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ClearScanner(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}