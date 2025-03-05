package main

import (
	"context"

	"eshop/kitex_gen/user"
	"log"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
)


// 登录
func loginHandler(c context.Context,ctx *app.RequestContext){
	logginReq := new(user.LoginReq)
	logginReq.Email = string(ctx.FormValue("email"))
	logginReq.Password = string(ctx.FormValue("password"))
	if userCli==nil{
		log.Fatal("userCli is nil")
	}
	resp,err:=userCli.Login(c,logginReq)
	// fmt.Println(logginReq.Email)
	// fmt.Println(logginReq.Password)
	if err!=nil{
		log.Printf("error occur in loginHandler,%v\n",err)
		// ctx.String(http.StatusBadRequest,"username or passord is invalid")
		ctx.JSON(http.StatusBadRequest,err)
	}else{
		// ctx.String(http.StatusOK,resp.Token)
		ctx.JSON(http.StatusOK,resp)
	}
}

// 注册
func registerHandler(c context.Context,ctx  *app.RequestContext){
	registerReq := new(user.RegisterReq)
	registerReq.Email = string(ctx.FormValue("email"))
	registerReq.ConfirmPassword = string(ctx.FormValue("confirm_password"))
	registerReq.Password = string(ctx.FormValue("password"))
	resp,err:=userCli.Register(c,registerReq)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,err)
	}else{
		ctx.JSON(http.StatusOK,resp)
	}

}
