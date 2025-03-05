package main

import (
	"context"
	"eshop/kitex_gen/cart"
	checkout "eshop/kitex_gen/checkout"
	"eshop/kitex_gen/order"
	"eshop/kitex_gen/payment"
	"fmt"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}


// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// TODO: Your code here...
	// orderReq:=new(order.PlaceOrderReq)

	// 获取购物车信息
	cartReq := new(cart.GetCartReq)
	cartReq.UserId = req.UserId
	cart,err := cartCli.GetCart(ctx,cartReq)
	if err!=nil{
		err = fmt.Errorf("fail to get the cart Info,%v",err)
		return
	}
	
	// 下单
	orderReq,amount := makeOrderReq(cart,req)
	orderResp,err := orderCli.PlaceOrder(ctx,orderReq)
	if err!=nil{
		err = fmt.Errorf("failed to place order,%v",err)
		return
	}
	// 支付

	payReq := new(payment.ChargeReq)
	payReq.Amount = amount
	payReq.CreditCard = req.CreditCard
	payReq.UserId = req.UserId
	payReq.OrderId = orderResp.Order.OrderId
	payResp,err :=payCli.Charge(ctx,payReq)

	if err!=nil{
		err=fmt.Errorf("failed to pay, %v",err)
		return
	}
	
	// 支付失败时，撤销订单
	// Todo

	resp = new(checkout.CheckoutResp)
	resp.OrderId = orderResp.Order.OrderId
	resp.TransactionId = payResp.TransactionId
	
	return
}

func makeOrderReq(cart *cart.GetCartResp,req *checkout.CheckoutReq)(*order.PlaceOrderReq,float32){
	orderReq := new(order.PlaceOrderReq)
	orderAdd := new(order.Address)
	orderAdd.City = req.Address.City
	orderAdd.Country = "china"
	orderAdd.State = req.Address.State
	orderAdd.StreetAddress = req.Address.StreetAddress
	orderReq.Address = orderAdd

	orderReq.Email = req.Email
	orderReq.UserCurrency = "cny"
	orderReq.UserId = req.UserId

	var amount float32
	for _,item:= range cart.Cart.Items{
		oItem := new(order.OrderItem)
		oItem.Item = item
		// cost需要调用product服务，暂时模拟成100，后续改写完接口时改进
		oItem.Cost = 100
		amount+=oItem.Cost
		orderReq.OrderItems = append(orderReq.OrderItems, oItem)
	}
	return orderReq,amount
}
