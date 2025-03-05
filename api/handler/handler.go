package handler

import (
	"context"
	// "eshop/api/config"
	"eshop/kitex_gen/user"
	"eshop/kitex_gen/user/userservice"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)


func login(ctx context.Context,c *app.RequestContext,userCli userservice.Client){
	req := new(user.LoginReq)
	if err:= c.BindAndValidate(&req);err!=nil{
		c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid login request"})
		return
	}
	userCli.Login(ctx,req)
}