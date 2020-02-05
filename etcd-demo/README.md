
### 服务端 

```go
// 连接etcd集群
cli, err := clientv3.New(clientv3.Config{
    // etcd集群成员节点列表
    Endpoints:   []string{"127.0.0.1:2379"},
    DialTimeout: 5 * time.Second,
})
if err != nil {
    fmt.Println("[测] connect etcd err:", err)
    return
}
// 创建命名解析
r := naming.GRPCResolver{Client: cli}
// 将本服务注册添加etcd中，服务名称为myService，服务地址为本机8001端口
r.Update(context.TODO(), "myService", naming.Update{Op: naming.Add, Addr: "192.168.0.101:8001"})
// TODO
// 服务在8001启动...
// …

```

### 客户端 

```go
// 连接etcd集群
cli, err := clientv3.New(clientv3.Config{
    // etcd集群成员节点列表
    Endpoints:   []string{"127.0.0.1:2379"},
    DialTimeout: 5 * time.Second,
})
if err != nil {
    fmt.Println("[测] connect etcd err:", err)
    return
}

r := &etcdnaming.GRPCResolver{Client: cli}
b := grpc.RoundRobin(r)

conn, err := grpc.Dial("myService", grpc.WithBalancer(b), grpc.WithBlock())
if err != nil {
    panic(err)
}

// TODO
// 使用conn创建具体服务的client，如userCli := protos.NewIUserServiceClient(conn)
```