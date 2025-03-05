package main

import (
	"context"
	"eshop/dao/mysql"
	cart "eshop/kitex_gen/cart"
	"eshop/model"
	"net/http"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// TODO: Your code here...
	println("cart additem")
	ncart := new(model.Cart)
	ncart.UserID = uint(req.UserId)
	ncart.ProductID = uint(req.Item.ProductId)
	ncart.Quantity = int(req.Item.Quantity)
	err =mysql.CreateCart(ncart)
	resp = new(cart.AddItemResp)
	if err!=nil{
		resp.Status = http.StatusBadRequest
	}else{
		resp.Status=200
	}
	return
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// TODO: Your code here...
	cartItems,err:=mysql.GetCartByID(req.UserId)
	resp = new(cart.GetCartResp)
	resp.Cart.UserId = req.UserId
	for _,item := range *cartItems{
		resp.Cart.Items = append(resp.Cart.Items, &cart.CartItem{Quantity: int32(item.Quantity),ProductId: uint32(item.ProductID)})
	}
	return
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// TODO: Your code here...
	resp = new(cart.EmptyCartResp)
	err = mysql.DeleteCartByID(req.UserId)
	if err!=nil{
		resp.Status=400
	}else{
		resp.Status=200
	}
	return
}
