package main

import (
	"eshop/dao/mysql"
	"eshop/kitex_gen/auth/authservice"
	user "eshop/kitex_gen/user/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	authCli authservice.Client
)

func main() {
	mysql.InitDB()
	// mysql.Init()
	// 服务注册
	r,err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err!=nil{
		log.Fatal(err)
	}
	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{IP:net.ParseIP("127.0.0.1"),Port: 9001}),
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "eshop.user",
			},
		),
	)
	// 服务发现
	resolver,err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err!=nil{
		log.Fatal(err)
	}
	c,err := authservice.NewClient("eshop.auth",client.WithResolver(resolver))
	if err!=nil{
		log.Fatal(err)
	}
	authCli = c
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
