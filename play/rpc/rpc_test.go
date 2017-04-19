// Copyright 2013 René Kistl. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package srpc

import (
	"errors"
	"fmt"
	//"github.com/pcdummy/skynet2/rpc/bsonrpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ugorji/go/codec"
)

var (
	serverBSON, serverGob, serverMsgPack, serverBinc, serverJSON string
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

// Some of Arith's methods have value args, some have pointer args. That's deliberate.

func (t *Arith) Add(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Div(args Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = args.A / args.B
	return nil
}

func (t *Arith) String(args *Args, reply *string) error {
	*reply = fmt.Sprintf("%d+%d=%d", args.A, args.B, args.A+args.B)
	return nil
}

func (t *Arith) Scan(args string, reply *Reply) (err error) {
	_, err = fmt.Sscan(args, &reply.C)
	return
}

func (t *Arith) Error(args *Args, reply *Reply) error {
	panic("ERROR")
}

func listenTCP() (net.Listener, string) {
	l, e := net.Listen("tcp", "127.0.0.1:0") // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

func TestStartServers(t *testing.T) {
	startBSONServer()
	startMsgPackServer()
	startGobServer()
	startBincServer()
	startJSONServer()
}

func startBSONServer() {
	rpc.Register(new(Arith))

	var l net.Listener
	l, serverBSON = listenTCP()
	// log.Println("Test BSON RPC server listening on", serverAddr)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err, conn)
			}
			// codec := bsonrpc.NewServerCodec(conn)
			// go rpc.ServeCodec(codec)
		}
	}()
}

func startMsgPackServer() {
	var l net.Listener
	l, serverMsgPack = listenTCP()
	// log.Println("Test MsgPack RPC server listening on", serverAddr2)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}
			codec := codec.GoRpc.ServerCodec(conn, &codec.MsgpackHandle{})
			go rpc.ServeCodec(codec)
		}
	}()
}

func startGobServer() {
	var l net.Listener
	l, serverGob = listenTCP()
	// log.Println("Test MsgPack RPC server listening on", serverAddr2)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}

			go rpc.ServeConn(conn)
		}
	}()
}

func startBincServer() {
	var l net.Listener
	l, serverBinc = listenTCP()
	// log.Println("Test MsgPack RPC server listening on", serverAddr2)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}

			codec := codec.GoRpc.ServerCodec(conn, &codec.BincHandle{})
			go rpc.ServeCodec(codec)
		}
	}()
}

func startJSONServer() {
	var l net.Listener
	l, serverJSON = listenTCP()
	// log.Println("Test MsgPack RPC server listening on", serverAddr2)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("accept error:", err)
			}

			codec := jsonrpc.NewServerCodec(conn)
			go rpc.ServeCodec(codec)
		}
	}()
}

func benchmarkClient(client *rpc.Client, b *testing.B) {
	// Synchronous calls
	args := &Args{7, 8}
	procs := runtime.GOMAXPROCS(-1)
	N := int32(b.N)
	var wg sync.WaitGroup
	wg.Add(procs)
	b.StartTimer()

	for p := 0; p < procs; p++ {
		go func() {
			reply := new(Reply)
			for atomic.AddInt32(&N, -1) >= 0 {
				err := client.Call("Arith.Add", args, reply)
				if err != nil {
					b.Fatalf("rpc error: Add: expected no error but got string %q", err.Error())
				}
				if reply.C != args.A+args.B {
					b.Fatalf("rpc error: Add: expected %d got %d", reply.C, args.A+args.B)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkEndToEndBSON(b *testing.B) {
	b.StopTimer()

	conn, err := net.Dial("tcp", serverBSON)
	if err != nil {
		log.Fatal("error dialing:", err, conn)
	}
	_ = conn

	// client := bsonrpc.NewClient(conn)
	// defer client.Close()

	// benchmarkClient(client, b)
}

func BenchmarkEndToEndMsgPack(b *testing.B) {
	b.StopTimer()

	conn, err := net.Dial("tcp", serverMsgPack)
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	rpcCodec := codec.GoRpc.ClientCodec(conn, &codec.MsgpackHandle{})
	client := rpc.NewClientWithCodec(rpcCodec)
	defer client.Close()

	benchmarkClient(client, b)
}

func BenchmarkEndToEndGob(b *testing.B) {
	b.StopTimer()

	conn, err := net.Dial("tcp", serverGob)
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	client := rpc.NewClient(conn)
	defer client.Close()

	benchmarkClient(client, b)
}

func BenchmarkEndToEndBinc(b *testing.B) {
	b.StopTimer()

	conn, err := net.Dial("tcp", serverBinc)
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	rpcCodec := codec.GoRpc.ClientCodec(conn, &codec.BincHandle{})
	client := rpc.NewClientWithCodec(rpcCodec)
	defer client.Close()

	benchmarkClient(client, b)
}

func BenchmarkEndToEndJSON(b *testing.B) {
	b.StopTimer()

	conn, err := net.Dial("tcp", serverJSON)
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	rpcCodec := jsonrpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(rpcCodec)
	defer client.Close()

	benchmarkClient(client, b)
}
