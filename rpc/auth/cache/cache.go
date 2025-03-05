package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStruct struct{
	client *redis.Client
}

func NewRedisStruct(add string)*RedisStruct{
	return &RedisStruct{
		client: redis.NewClient(
			&redis.Options{
				Addr: add,
			},
		),
	}
}

func (c *RedisStruct)AddToken(token string, userID uint32)error{
	key:=fmt.Sprintf("token:%s",token)
	return c.client.Set(context.Background(),key,userID,time.Minute*10).Err()
}


func (c *RedisStruct)VerifyToken(key string)(uint32,error){
	key,err :=c.client.Get(context.Background(),key).Result()
	if err!=nil{
		return 0,err
	}
	val,err :=strconv.Atoi(key)
	return uint32(val),err
}