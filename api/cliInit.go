package main

import (
	"eshop/kitex_gen/cart/cartservice"
	"eshop/kitex_gen/order/orderservice"
	"eshop/kitex_gen/payment/paymentservice"
	"eshop/kitex_gen/product/productcatalogservice"
	"eshop/kitex_gen/user/userservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	// "google.golang.org/protobuf/internal/order"
)


var (
	userCli userservice.Client
	productCli productcatalogservice.Client
	orderCli orderservice.Client
	cartCli cartservice.Client
	payCli paymentservice.Client
)

func cliInit(){
	r,err:=etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err!=nil{
		log.Fatal(err)
	}

	// 初始化user客户端
	u ,err := userservice.NewClient("eshop.user",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	userCli = u

	// 初始化product客户端
	pro,err := productcatalogservice.NewClient("eshop.product",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	productCli = pro

	// 初始化order客户端
	ord,err := orderservice.NewClient("eshop.order",client.WithResolver(r))
	if err != nil{
		log.Fatal(err)
	}
	orderCli = ord

	// 初始化cart客户端
	cart ,err := cartservice.NewClient("eshop.cart",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	cartCli = cart

	// 初始化payment客户端
	pay,err := paymentservice.NewClient("eshop.payment",client.WithResolver(r))
	if err !=nil{
		log.Fatal(err)
	}
	payCli = pay
}