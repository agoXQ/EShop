// package main

// import (
// 	auth "eshop/kitex_gen/auth/authservice"
// 	"log"
// 	"net"

// 	"github.com/cloudwego/kitex/pkg/rpcinfo"
// 	"github.com/cloudwego/kitex/server"
// 	etcd "github.com/kitex-contrib/registry-etcd"
// )

// func main() {

// 	// 消息队列相关操作
	

// 	// 服务注册
// 	r,err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	svr := auth.NewServer(
// 		new(AuthServiceImpl),
// 		server.WithServiceAddr(&net.TCPAddr{IP:net.ParseIP("127.0.0.1"),Port: 9000}),
// 		server.WithRegistry(r),
// 		server.WithServerBasicInfo(
// 			&rpcinfo.EndpointBasicInfo{
// 				ServiceName: "eshop.auth",
// 			},
// 		),
// 	)

// 	err = svr.Run()

// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// }


package main

import (
    "log"
    "net"
    "sync"

    auth "eshop/kitex_gen/auth/authservice"
    "github.com/cloudwego/kitex/pkg/rpcinfo"
    "github.com/cloudwego/kitex/server"
    etcd "github.com/kitex-contrib/registry-etcd"
	"eshop/rpc/auth/rabbitmq"
)

func main() {
    // 使用 WaitGroup 等待所有 goroutine 完成
    var wg sync.WaitGroup
    wg.Add(2) // 两个任务：Kitex 服务器和 RabbitMQ 消费者

    // 启动 RabbitMQ 消费者
    go func() {
        defer wg.Done()
        rabbitmq.StartTokenVerificationConsumer()
    }()

    // 启动 Kitex 服务器
    go func() {
        defer wg.Done()

        // 服务注册
        r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
        if err != nil {
            log.Fatal(err)
        }

        svr := auth.NewServer(
            new(AuthServiceImpl),
            server.WithServiceAddr(&net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9000}),
            server.WithRegistry(r),
            server.WithServerBasicInfo(
                &rpcinfo.EndpointBasicInfo{
                    ServiceName: "eshop.auth",
                },
            ),
        )

        err = svr.Run()
        if err != nil {
            log.Println(err.Error())
        }
    }()

    // 等待所有 goroutine 完成
    wg.Wait()
}
