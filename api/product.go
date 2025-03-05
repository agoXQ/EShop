package main

import (
	"context"
	"eshop/kitex_gen/product"
	"log"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func ListProductsHandler(ctx context.Context,rctx *app.RequestContext){
	req := new(product.ListProductsReq)
	if err:=rctx.BindJSON(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid request body"})
		return
	}
	
	resp,err:=productCli.ListProducts(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":err})
		return
	}
	log.Printf("oc")
	log.Printf("%v",resp)
	rctx.JSON(http.StatusOK,resp)
}

func GetProductHandler(ctx context.Context,rctx *app.RequestContext){
	req := new(product.GetProductReq)
	if err:=rctx.Bind(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid request body"})
		return 
	}
	resp,err:=productCli.GetProduct(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":err})
		return
	}
	rctx.JSON(http.StatusOK,resp)
}	

func SearchProductsHandler(ctx context.Context,rctx *app.RequestContext){
	req:=new(product.SearchProductsReq)
	if err:=rctx.Bind(req);err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":"Invalid request body"})
		return
	}

	resp,err := productCli.SearchProducts(ctx,req)
	if err!=nil{
		rctx.JSON(http.StatusBadRequest,utils.H{"error":err})
		return
	}

	rctx.JSON(http.StatusOK,resp)
}