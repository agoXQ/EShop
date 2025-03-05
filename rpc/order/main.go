package main

import (
	"eshop/dao/mysql"
	order "eshop/kitex_gen/order/orderservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	// 初始化DB
	if mysql.DB == nil {
		mysql.InitDB()
	}

	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	svr := order.NewServer(
		new(OrderServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9003}),
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "eshop.order",
			},
		),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
