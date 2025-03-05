package service

import (
	"eshop/rpc/auth/cache"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

type ServiceStruct struct{
	cache *cache.RedisStruct
}

func NewServiceStruct(addr string)*ServiceStruct{
	return &ServiceStruct{
		cache: cache.NewRedisStruct(addr),
	}
}

func (s *ServiceStruct)Verify(t string)(bool,error){
	// 缓存
	key:= fmt.Sprintf("token:%s",t)
	if _,err:=s.cache.VerifyToken(key);err==nil{
		return true,nil
	}

	// 1.解析token
	token,err := jwt.Parse(t,func(t *jwt.Token) (interface{}, error) {
		return []byte("abcdmfg"),nil
	})
	if err!=nil{
		// s.cache.AddToken(token,)
		return false,err
		// return resp,nil
	}


	// 2.验证token是否有效
	if claims,ok:=token.Claims.(jwt.MapClaims);ok&&token.Valid{
		userID,ok:=claims["user_id"].(float64)
		if ok{
			s.cache.AddToken(t,uint32(userID))
		}
		log.Printf("token验证通过，用户id：%v ",claims["user_id"])
		// s.cache.AddToken(t,claims["user_id"].uint32())
		return true,nil
	}else{
		log.Printf("token验证未通过，用户id：%v ",claims["user_id"])
		return false,err
	}
	// return
}