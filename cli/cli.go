package main

import (
	"context"
	"demo/gothrift/service/echo"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	hostPort := net.JoinHostPort("127.0.0.1", "9898")
	transport, err := thrift.NewTSocketTimeout(hostPort, time.Second*10)
	if err != nil {
		log.Panic("error resolving address", err)
	}

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		log.Panic("resolving address", err)
	}

	if err = transport.Open(); err != nil {
		log.Panic("error opening socket to 127.0.0.1:9898", err)
	}
	defer func() {
		_ = transport.Close()
	}()

	go func() {
		client := echo.NewEchoClientFactory(useTransport, protocolFactory)
		for i := 0; i < 1000000; i++ {
			var resp *echo.EchoResponse
			req := &echo.EchoRequest{
				Msg: "You are welcome go thrift.",
				Tag: int32(i),
			}
			resp, err = client.Echo(context.Background(), req)
			if err != nil {
				log.Println("Echo failed:", err)
				return
			}
			log.Printf("response[%d]: %v", resp.GetTag(), resp.Msg)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	log.Println(<-c)
}
