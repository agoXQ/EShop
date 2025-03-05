package cache

import (
	"context"
	"encoding/json"
	"eshop/model"
	"time"

	"github.com/go-redis/redis/v8"
)
type RedisStruct struct{
	client *redis.Client
}

func NewRedisStruct(Addr string ) (*RedisStruct){
	return &RedisStruct{
		client: redis.NewClient(&redis.Options{
		Addr: Addr,
		PoolSize: 10,              // 连接池大小
        MinIdleConns: 5,           // 最小空闲连接数
        ReadTimeout:  time.Second,  // 读取超时时间
        WriteTimeout: time.Second,  // 写入超时时间
        DialTimeout:  time.Second,  // 建立连接超时时间
	}),
	}
}

func (c *RedisStruct)GetProduct(key string)(*model.Product,error){
	val,err := c.client.Get(context.Background(),key).Result()
	if err!=nil{
		return nil,err
	}
	var product model.Product
	if err:=json.Unmarshal([]byte(val),&product);err!=nil{
		return nil,err
	}
	return &product,nil
}

func (c *RedisStruct) SetProduct(key string,product *model.Product)error{
	val,err:=json.Marshal(product)
	if err!=nil{
		return err
	}
	return c.client.Set(context.Background(),key,val,10*time.Second).Err()
}

func (c *RedisStruct) GetProducts(key string)([]model.Product,error){
	val,err := c.client.Get(context.Background(),key).Result()
	if err !=nil{
		return nil,err
	}

	var products []model.Product
	if err := json.Unmarshal([]byte(val),&products);err!=nil{
		return nil,err
	}
	return products,err
}

func (c *RedisStruct) SetProducts(key string,products []model.Product)error{
	val,err := json.Marshal(products)
	if err!=nil{
		return err
	}

	return c.client.Set(context.Background(),key,val,10*time.Second).Err()
}