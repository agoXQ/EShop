package model

import (
    "gorm.io/gorm"
    // "time"
)

type Transaction struct {
    gorm.Model
    TransactionID   string  `gorm:"primaryKey;column:transaction_id;type:varchar(36);not null"` // 交易ID
    OrderID         string  `gorm:"column:order_id;type:varchar(36);not null"`                // 订单ID
    UserID          uint    `gorm:"column:user_id;not null"`                                   // 用户ID
    Amount          float64 `gorm:"column:amount;type:decimal(10,2);not null"`                 // 支付金额
    CreditCardNumber string `gorm:"column:credit_card_number;type:varchar(20);not null"`       // 信用卡号（仅存储后4位）
    Status          string  `gorm:"column:status;type:varchar(20);not null"`                   // 支付状态（SUCCESS/FAILED）
}

func (t *Transaction) TableName() string {
    return "transactions"
}