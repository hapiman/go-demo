package thrift_demo

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go-demo/thrift-demo/gen-go/echo"
)

type EchoServerImp struct {
}

func (e *EchoServerImp) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())

	res := &echo.EchoRes{
		Msg: req.GetMsg(),
	}

	return res, nil
}

func SStart() {
	transport, err := thrift.NewTServerSocket(fmt.Sprintf(":%d", rpcPort))
	if err != nil {
		panic(err)
	}

	processor := echo.NewEchoProcessor(&EchoServerImp{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(bufferSize),
		thrift.NewTCompactProtocolFactory(),
	)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
