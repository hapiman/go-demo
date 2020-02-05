package thrift_demo

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go-demo/thrift-demo/gen-go/echo"
	"os"
	"strings"
)

func CStart() {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(fmt.Sprintf("localhost:%d", rpcPort))
	if err != nil {
		fmt.Errorf("NewTSocket failed. err: [%v]\n", err)
		return
	}

	transport, err = thrift.NewTBufferedTransportFactory(bufferSize).GetTransport(transport)
	if err != nil {
		fmt.Errorf("NewTransport failed. err: [%v]\n", err)
		return
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		fmt.Errorf("Transport.Open failed. err: [%v]\n", err)
		return
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	// 生成client
	client := echo.NewEchoClient(thrift.NewTStandardClient(iprot, oprot))

	var res *echo.EchoRes
	res, err = client.Echo(context.Background(), &echo.EchoReq{
		Msg: strings.Join(os.Args[1:], " "),
	})
	if err != nil {
		fmt.Errorf("client echo failed. err: [%v]", err)
		return
	}

	fmt.Printf("message from server: %v", res.GetMsg())
}
