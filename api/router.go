package main

import "github.com/cloudwego/hertz/pkg/app/server"

func Register()(*server.Hertz,error){
	r:=server.New(server.WithHostPorts("localhost:8888"))
	auth := r.Group("/auth")
	auth.POST("/verify",verifyHandler)

	user := r.Group("/user")
	user.POST("/login",	loginHandler)
	user.POST("/register",registerHandler)

	product:=r.Group("/product")
	product.POST("/list",ListProductsHandler)
	product.POST("/get",GetProductHandler)
	product.POST("/search",SearchProductsHandler)

	order := r.Group("/order")
	order.POST("/place",placeOrderHandler)
	// order.POST("")

	cart := r.Group("/cart")
	cart.POST("/get",getCartHandler)
	cart.POST("/empty",emptyCartHandler)
	cart.POST("/add",addItemHandler)

	payment := r.Group("/payment")
	payment.POST("/charge",charge)

	return r,nil
}