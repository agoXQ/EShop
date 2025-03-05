package main

import (
	"eshop/dao/mysql"
	"eshop/kitex_gen/cart/cartservice"
	checkout "eshop/kitex_gen/checkout/checkoutservice"
	"eshop/kitex_gen/order/orderservice"
	"eshop/kitex_gen/payment/paymentservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	// amqp "github.com/rabbitmq/amqp091-go"
)

var (
	payCli paymentservice.Client
	orderCli orderservice.Client
	cartCli cartservice.Client
)

func CliInit(){
	r,err:=etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err!=nil{
		log.Fatal(err)
	}

	// 初始化支付客户端
	pay,err:=paymentservice.NewClient("eshop.pay",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	payCli = pay

	// 初始化订单客户端
	order,err := orderservice.NewClient("eshop.order",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	orderCli = order

	// 初始化购物车客户端
	cart,err := cartservice.NewClient("eshop.cart",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	cartCli = cart
}

func main() {
	CliInit()

	// svr := checkout.NewServer(new(CheckoutServiceImpl))
	// // 9005
	// err := svr.Run()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// 初始化DB
	if mysql.DB == nil {
		mysql.InitDB()
	}

	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	svr := checkout.NewServer(
		new(CheckoutServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9005}),
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

	// conn,err := amqp.Dial()
}


