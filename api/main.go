package main

import (
	// "eshop/api/config"
	// "eshop/kitex_gen/user/userservice"
	"log"

	// "github.com/cloudwego/kitex/client"
	// "github.com/cloudwego/hertz/pkg/app"
	// "github.com/cloudwego/hertz/pkg/app/server"
	// etcd "github.com/kitex-contrib/registry-etcd"
)



func main(){
	// r,err :=etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	
	// // userCli,err = userservice.NewClient("eshop.user",client.WithResolver(r))
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	cliInit()
	eshopServer,err:=Register()
	if err!=nil{
		log.Fatal(err)
	}
	eshopServer.Spin()
	// eshopServer.GET("/user/register",)
}