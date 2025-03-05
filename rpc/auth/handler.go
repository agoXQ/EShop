package main

import (
	"context"
	auth "eshop/kitex_gen/auth"
	"eshop/rpc/auth/cache"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	// "rpc/auth/cache"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

var r *cache.RedisStruct
func init(){
	r = cache.NewRedisStruct("localhost:6379")
}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// TODO: Your code here...
	// 使用jwt完成token生成
	// 1.生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id": req.UserId,
	})

	// 2.给token签名
	token_string,err :=token.SignedString([]byte("abcdmfg"))

	// 3.返回token
	r.AddToken(token_string,uint32(req.UserId))
	resp = new(auth.DeliveryResp)
	resp.Token = token_string
	return
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	//   在缓存中查询
	key:= fmt.Sprintf("token:%s",req.Token)
	if _,err := r.VerifyToken(key);err==nil{
		resp = new(auth.VerifyResp)
		resp.Res=true
		return resp,nil
	}
	
	// 1.解析token
	token,err := jwt.Parse(req.Token,func(t *jwt.Token) (interface{}, error) {
		return []byte("abcdmfg"),nil
	})
	if err!=nil{
		resp = &auth.VerifyResp{Res: false}
		return resp,nil
	}


	// 2.验证token是否有效
	if claims,ok:=token.Claims.(jwt.MapClaims);ok&&token.Valid{
		log.Printf("token验证通过，用户id：%v ",claims["user_id"])
		resp = &auth.VerifyResp{Res: true}
	}else{
		resp = &auth.VerifyResp{Res: false}
	}
	return
}
