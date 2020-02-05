package main

import (
	"context"
	"demo/gothrift/echo"
	"log"

	// "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/apache/thrift/lib/go/thrift"
)

type EchoServerImp struct {
}

func (e *EchoServerImp) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("message from client[%d]: %v\n", req.GetTag(), req.GetMsg())
	resp := &echo.EchoResponse{
		Msg: req.GetMsg(),
		Tag: req.GetTag(),
	}
	return resp, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	processor := echo.NewEchoProcessor(&EchoServerImp{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)

	log.Panic(server.Serve())
}
