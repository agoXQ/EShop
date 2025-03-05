package mysql

import (
	"errors"
	"eshop/model"
	"fmt"

	"gorm.io/gorm"
)

func CreateTransaction(tran *model.Transaction)error{
	return DB.Create(&tran).Error
}

func GetTransactionByOrderID(orderID string)(*model.Transaction,error){
	var transaction model.Transaction
    if err := DB.Where("order_id = ?", orderID).First(&transaction).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("transaction not found for order ID: %s", orderID)
        }
        return nil, fmt.Errorf("failed to query transaction: %v", err)
    }
	return &transaction,nil
}