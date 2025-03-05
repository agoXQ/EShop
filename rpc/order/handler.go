package main

import (
	"context"
	"eshop/dao/mysql"
	"eshop/kitex_gen/cart"
	order "eshop/kitex_gen/order"
	"eshop/model"
	"fmt"
	"time"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// TODO: Your code here...
	// 创建订单
	newOrder := new(model.Order)
	newOrder.OrderID = generateOrderID()
	newOrder.StreetAddress = req.Address.StreetAddress
	newOrder.City = req.Address.City
	newOrder.Country = req.Address.Country
	newOrder.State = req.Address.State

	newOrder.Email = req.Email
	for _, item := range req.OrderItems {
		orderitem := new(model.OrderItem)
		orderitem.ItemID = item.Item.ProductId
		orderitem.ItemQuantity = item.Item.Quantity
		product, _ := mysql.GetProductByID(item.Item.ProductId)
		orderitem.ItemName = product.Name
		orderitem.ItemCost = product.Price * float32(item.Item.Quantity)

		newOrder.OrderItems = append(newOrder.OrderItems, *orderitem)
	}

	err = mysql.CreateOrder(newOrder)
	resp = new(order.PlaceOrderResp)
	resp.Order = new(order.OrderResult)
	resp.Order.OrderId = newOrder.OrderID
	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// TODO: Your code here...
	// 查询订单
	orders, err := mysql.GetOrderByUserID(uint(req.UserId))
	if err != nil {
		resp.Orders = nil
	} else {
		for _, order := range orders {
			resp.Orders = append(resp.Orders, convertToProtoOrder(&order))
		}
	}

	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	// 标记支付信息
	 _ , err = mysql.GetTransactionByOrderID(req.OrderId)
	if err!=nil{
		return nil, fmt.Errorf("No transaction found,%v",err)
	}
	// resp = new(order.MarkOrderPaidResp)
	
	return
}

// 生成订单ID
func generateOrderID() string {
	return fmt.Sprintf("ORDER-%d", time.Now().UnixNano())
}

// 生成商品ID
func generateOrderItemID() string {
	return fmt.Sprintf("ITEM-%d", time.Now().UnixNano())
}

// 将数据库模型转换为 gRPC 消息
func convertToProtoOrder(morder *model.Order) *order.Order {

	var orderItems []*order.OrderItem
	for _, item := range morder.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: item.ItemID,
				// Name:     item.ItemName,
				Quantity: int32(item.ItemQuantity),
			},
			Cost: float32(item.ItemCost),
		})
	}
	return &order.Order{
		OrderItems:   orderItems,
		OrderId:      morder.OrderID,
		UserId:       uint32(morder.UserID),
		UserCurrency: morder.UserCurrency,
		Address: &order.Address{
			StreetAddress: morder.StreetAddress,
			City:          morder.City,
			State:         morder.State,
			Country:       morder.Country,
			ZipCode:       int32(morder.ZipCode),
		},
		Email:     morder.Email,
		CreatedAt: int32(morder.CreatedAt),
	}
}
