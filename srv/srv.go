package main

import (
	"context"
	"demo/gothrift/service/echo"
	"demo/gothrift/service/tapi"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
)

type EchoServerImp struct {
}

func (serve *EchoServerImp) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("message from client[%d]: %v\n", req.GetTag(), req.GetMsg())
	resp := &echo.EchoResponse{
		Msg: req.GetMsg(),
		Tag: req.GetTag(),
	}
	return resp, nil
}

type TapiServerImp struct {
}

func (serve *TapiServerImp) Regist(ctx context.Context, req *tapi.ReqRegist) (*tapi.AckRegist, error) {
	log.Printf("message from client %v", req)
	resp := &tapi.AckRegist{
		Code:    0,
		Message: "注册成功",
		Account: req.GetAccount(),
	}
	return resp, nil
}

func (serve *TapiServerImp) Login(ctx context.Context, req *tapi.ReqLogin) (*tapi.AckLogin, error) {
	log.Printf("message from client %v", req)
	resp := &tapi.AckLogin{
		Code:    0,
		Message: "登录成功",
	}
	return resp, nil
}

// func main() {
// 	transport, err := thrift.NewTServerSocket(":9898")
// 	if err != nil {
// 		log.Panic(err)
// 	}
//
// 	// protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
// 	// transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
//
// 	processor := echo.NewEchoProcessor(&EchoServerImp{})
// 	server := thrift.NewTSimpleServer4(
// 		processor,
// 		transport,
// 		thrift.NewTBufferedTransportFactory(8192),
// 		thrift.NewTCompactProtocolFactory(),
// 	)
//
// 	err = server.Serve()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }

func main() {
	serverTransport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		log.Panic(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	echoProcessor := echo.NewEchoProcessor(&EchoServerImp{})
	tapiProcessor := tapi.NewTapiProcessor(&TapiServerImp{})

	multiProcessor := thrift.NewTMultiplexedProcessor()
	// 给每个service起一个名字
	multiProcessor.RegisterProcessor("echo", echoProcessor)
	multiProcessor.RegisterProcessor("tapi", tapiProcessor)

	server := thrift.NewTSimpleServer4(
		multiProcessor,
		serverTransport,
		transportFactory,
		protocolFactory,
	)
	defer func() { _ = server.Stop() }()

	err = server.Serve()
	if err != nil {
		log.Panic(err)
	}
}
