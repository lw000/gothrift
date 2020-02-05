package main

import (
	"context"
	"demo/gothrift/echo"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
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
		log.Panic("error resolving address", err)
	}

	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err = transport.Open(); err != nil {
		log.Panic("error opening socket to 127.0.0.1:9898", err)
	}
	defer func() {
		_ = transport.Close()
	}()

	for i := 0; i < 10; i++ {
		var resp *echo.EchoResponse
		req := &echo.EchoRequest{Msg: "You are welcome.", Tag: int32(i)}
		resp, err = client.Echo(context.Background(), req)
		if err != nil {
			log.Println("Echo failed:", err)
			return
		}

		log.Printf("response[%d]: %v", resp.GetTag(), resp.Msg)
	}

	log.Println("well done")
}
