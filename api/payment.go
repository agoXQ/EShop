package main

import (
	"context"
	"eshop/kitex_gen/payment"
	"fmt"
	"net/http"

	// "unicode/utf16"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func charge(ctx context.Context,rctx *app.RequestContext){
	println("payment charge")
	req := new(payment.ChargeReq)
	if err:= rctx.BindJSON(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid Request Body"})
		return
	}
	resp,err:=payCli.Charge(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadGateway,utils.H{"error":fmt.Sprintf("Fail to pay the order,%v",err)})
		return
	}
	rctx.JSON(http.StatusOK,resp)	
}