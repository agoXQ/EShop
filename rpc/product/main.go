package main

import (
	"encoding/json"
	"eshop/dao/mysql"
	product "eshop/kitex_gen/product/productcatalogservice"
	"eshop/model"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var defaultPageSize = 100

func main() {
	// 初始化时给数据库添加一些东西
	
	if 	mysql.DB ==nil{
		mysql.InitDB()
	}
	add()
	
	r,err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err!=nil{
		log.Fatal(err)
	}
	svr := product.NewServer(
		new(ProductCatalogServiceImpl),
		server.WithServiceAddr(&net.TCPAddr{IP: net.ParseIP("127.0.0.1"),Port: 9002}),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName:"eshop.product",
		}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func add(){
	products := []*model.Product{}
	names := []string{"abc","def","ghi","adj","adv"}
	for i:=0;i<5;i++{
		pro:=model.Product{}
		categories :=[]string{"car","food","clothes"}
		categoriestring,_ :=json.Marshal(categories)
		pro.Categories = string(categoriestring)
		pro.Name = names[i]
		pro.Picture = "/picture/"+names[i]+".png"
		pro.Price = float32(12*i+50.0)
		products = append(products, &pro)
	}
	if err :=mysql.CreateProduct(products);err!=nil{
		log.Fatal(err)
	}
}
