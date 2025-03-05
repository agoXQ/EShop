package main

import (
	"context"
	"eshop/dao/mysql"
	"eshop/kitex_gen/auth"
	user "eshop/kitex_gen/user"
	"eshop/model"
	"fmt"
	"log"
	// "github.com/pingcap/tidb/parser/mysql"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterResp)
	u:=model.User{
		Email: req.Email,
		Password: req.Password,
	}
	if err:=mysql.CreateUser([]*model.User{&u});err!=nil{
		resp.UserId = 0
	}
	resp.UserId,err = mysql.GetUserID(u.Email)
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	resp = new(user.LoginResp)
	if user,err := mysql.GetUser(req.Email);err !=nil{
		return nil,err
	}else{
		if user.Password==req.Password{
			resp.UserId = int32(user.ID)
			authReq := new(auth.DeliverTokenReq)
			authReq.UserId = resp.UserId
			authResp,err := authCli.DeliverTokenByRPC(ctx,authReq)
			if err!=nil{
				log.Printf("token分发失败")
				return nil,err
			}else{
				resp.Token = authResp.Token
			}
		}else{
			return nil,fmt.Errorf("邮箱或者密码错误，请重新检查后输入")
		}
	}
	
	return
}
