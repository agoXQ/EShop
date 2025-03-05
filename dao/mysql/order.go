package mysql

import (
	"errors"
	// "eshop/kitex_gen/order"
	"eshop/model"

	"gorm.io/gorm"
)

func CreateOrder(order *model.Order)error{
	// 开启事务
	tx:=DB.Begin()
	tx.Exec("SET TRANSACTION ISOLATION LEVEL READ COMMITTED")
	defer func(){
		if r:=recover();r!=nil{
			tx.Rollback()
		}
	}()

	// 插入订单
	if err:=tx.Create(order).Error;err!=nil{
		tx.Rollback()
		return errors.New("fail to create order")
	}

	// 插入订单项
	for _,item := range order.OrderItems{
		item.OrderID = order.ID // 确保order ID的准确性
		if err := tx.Create(item);err!=nil{
			tx.Rollback()
			return errors.New("fail to create item")
		}
	}

	// 提交事务
	if err:= tx.Commit();err!=nil{
		return errors.New("fail to commit transaction")
	}
	return nil

}

// 根据userID查询订单ID
func GetOrderByUserID(userID uint)([]model.Order,error){
	var orders = []model.Order{}
	if err:= DB.Preload("OrderItems").Where("user_id = ?",userID).Find(&orders).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("order not found")
        }
        return nil, errors.New("failed to query order")
		// return nil,err
	}
	return orders,nil
}

// GetOrderByID 根据订单ID查询订单及其订单项
func GetOrderByID(orderID string) (*model.Order, error) {
    var order model.Order
    if err := DB.Preload("OrderItems").Where("order_id = ?", orderID).First(&order).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("order not found")
        }
        return nil, errors.New("failed to query order")
    }
    return &order, nil
}