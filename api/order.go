package main

import (
	"context"
	"eshop/kitex_gen/order"
	// "fmt"
	// "io/ioutil"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func placeOrderHandler(ctx context.Context,rctx *app.RequestContext){
	println("order handler")
    req := new(order.PlaceOrderReq)
    if err := rctx.BindJSON(req); err != nil {
        rctx.JSON(http.StatusBadRequest, utils.H{"error": "Invalid Request Body"})
        return
    }
	resp,err :=orderCli.PlaceOrder(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadGateway,utils.H{"error":err})
		return 
	}
	rctx.JSON(http.StatusOK,resp)
}

// func placeOrderHandler(ctx context.Context, rctx *app.RequestContext) {
//     body, err := ioutil.ReadAll(rctx.Request.Body)
//     if err != nil {
//         rctx.JSON(http.StatusBadRequest, utils.H{"error": "Failed to read request body"})
//         return
//     }
//     fmt.Println("Received body:", string(body)) // 打印请求体内容

//     req := new(order.PlaceOrderReq)
//     if err := rctx.BindJSON(req); err != nil {
//         rctx.JSON(http.StatusBadRequest, utils.H{"error": "Invalid Request Body"})
//         return
//     }
//     // 继续处理逻辑...
// }