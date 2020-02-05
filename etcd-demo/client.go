package etcd_demo

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)
func Start (){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	resp, err := cli.Put(ctx, "/sample_key", "sample_value")
	cancel()
	if err != nil {
		panic(err)
	}
	// use the response
	fmt.Print(resp)
}