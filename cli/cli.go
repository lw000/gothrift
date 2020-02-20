package main

import (
	"context"
	"demo/gothrift/rpc/echo"
	"demo/gothrift/rpc/tapi"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

//
// func main() {
// 	hostPort := net.JoinHostPort("127.0.0.1", "9898")
// 	transport, err := thrift.NewTSocketTimeout(hostPort, time.Second*10)
// 	if err != nil {
// 		log.Panic("error resolving address", err)
// 	}
//
// 	transportFactory := thrift.NewTBufferedTransportFactory(8192)
// 	protocolFactory := thrift.NewTCompactProtocolFactory()
// 	useTransport, err := transportFactory.GetTransport(transport)
// 	if err != nil {
// 		log.Panic("resolving address", err)
// 	}
//
// 	if err = transport.Open(); err != nil {
// 		log.Panic("error opening socket to 127.0.0.1:9898", err)
// 	}
// 	defer func() {
// 		_ = transport.Close()
// 	}()
//
// 	go func() {
// 		client := echo.NewEchoClientFactory(useTransport, protocolFactory)
// 		for i := 0; i < 1000000; i++ {
// 			var resp *echo.EchoResponse
// 			req := &echo.EchoRequest{
// 				Msg: "You are welcome go thrift.",
// 				Tag: int32(i),
// 			}
// 			resp, err = client.Echo(context.Background(), req)
// 			if err != nil {
// 				log.Println("Echo failed:", err)
// 				return
// 			}
// 			log.Printf("response[%d]: %v", resp.GetTag(), resp.Msg)
// 		}
// 	}()
//
// 	c := make(chan os.Signal)
// 	signal.Notify(c, os.Interrupt)
//
// 	log.Println(<-c)
// }

func runEcho() {
	hostPort := net.JoinHostPort("127.0.0.1", "9898")
	socket, err := thrift.NewTSocketTimeout(hostPort, time.Second*10)
	if err != nil {
		log.Panic("error resolving address", err)
	}
	transport := thrift.NewTFramedTransport(socket)
	protocol := thrift.NewTBinaryProtocolTransport(transport)

	if err = transport.Open(); err != nil {
		log.Panic("error opening socket to 127.0.0.1:9898", err)
	}
	defer func() {
		_ = transport.Close()
	}()

	ctx := context.Background()
	echoProtocol := thrift.NewTMultiplexedProtocol(protocol, "echo")
	standardClient := thrift.NewTStandardClient(echoProtocol, echoProtocol)
	client := echo.NewEchoClient(standardClient)
	for i := 0; i < 1; i++ {
		var resp *echo.EchoResponse
		req := &echo.EchoRequest{
			Msg: "You are welcome go thrift.",
			Tag: int32(i),
		}
		resp, err = client.Echo(ctx, req)
		if err != nil {
			log.Println("Echo failed:", err)
			return
		}
		log.Printf("response[%d]: %v", resp.GetTag(), resp.Msg)
	}
}

func runTapi() {
	hostPort := net.JoinHostPort("127.0.0.1", "9898")
	socket, err := thrift.NewTSocketTimeout(hostPort, time.Second*10)
	if err != nil {
		log.Panic("error resolving address", err)
	}
	transport := thrift.NewTFramedTransport(socket)
	protocol := thrift.NewTBinaryProtocolTransport(transport)

	if err = transport.Open(); err != nil {
		log.Panic("error opening socket to 127.0.0.1:9898", err)
	}
	defer func() {
		_ = transport.Close()
	}()

	ctx := context.Background()
	tapiProtocol := thrift.NewTMultiplexedProtocol(protocol, "tapi")
	standardClient := thrift.NewTStandardClient(tapiProtocol, tapiProtocol)
	client := tapi.NewTapiClient(standardClient)

	{
		var resp *tapi.AckRegist
		req := &tapi.ReqRegist{
			Account:  "lw00000",
			Password: "111111",
		}
		resp, err = client.Regist(ctx, req)
		if err != nil {
			log.Println("tapi failed:", err)
			return
		}
		log.Printf("response %v", resp)
	}

	{
		var resp *tapi.AckLogin
		req := &tapi.ReqLogin{
			Account:  "lw00000",
			Password: "111111",
		}
		resp, err = client.Login(ctx, req)
		if err != nil {
			log.Println("tapi failed:", err)
			return
		}
		log.Printf("response %v", resp)
	}
}

func main() {
	for i := 0; i < 1; i++ {
		go runEcho()
	}

	go runTapi()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	log.Println(<-c)
}
