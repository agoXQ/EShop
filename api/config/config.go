package config

import (
	"eshop/kitex_gen/user/userservice"
	"log"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type Config struct{
	userCli userservice.Client
}

func NewConfig() (*Config,error){
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil{
		log.Fatal(err)
	}
	cfg := new(Config)
	cfg.userCli,err = userservice.NewClient("eshop.user",client.WithResolver(r))
	if err!=nil{
		log.Fatal(err)
	}
	return cfg,nil
}