package main

import (
	"context"
	"eshop/kitex_gen/cart"
	"net/http"

	// "eshop/kitex_gen/cart/cartservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func addItemHandler(ctx context.Context,rctx *app.RequestContext){
	// 往购物车添加商品
	println("cart add")
	req := new(cart.AddItemReq)
	if err:=rctx.BindJSON(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid Request Body"})
		return
	}
	resp,err:=cartCli.AddItem(ctx,req)
	if err!=nil{
		// println(resp.Status)
		rctx.JSON(http.StatusBadGateway,utils.H{"error":"fail to add cart item, please try again soon"})
		return
	}
	rctx.JSON(http.StatusOK,resp)

}

func emptyCartHandler(ctx context.Context,rctx *app.RequestContext){
	// 清空购物车
	req := new(cart.EmptyCartReq)
	if err:=rctx.BindJSON(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid Request Body"})
		return
	}

	resp,err:=cartCli.EmptyCart(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadGateway,utils.H{"error":err})
		return
	}
	rctx.JSON(http.StatusOK,resp)
}

func getCartHandler(ctx context.Context,rctx *app.RequestContext){
	// 获取购物车

	req := new(cart.GetCartReq)
	if err:=rctx.BindJSON(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid Request Body"})
		return
	}

	resp,err:=cartCli.GetCart(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadGateway,utils.H{"error":err})
		return
	}
	rctx.JSON(http.StatusOK,resp)
}