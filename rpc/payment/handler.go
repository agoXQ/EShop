package main

import (
	"context"
	"errors"
	"eshop/dao/mysql"
	payment "eshop/kitex_gen/payment"
	"eshop/model"
	"fmt"
	"time"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// TODO: Your code here...
	transactionID,err := processPayment(req)
	if err!=nil{
		return nil,fmt.Errorf("payment failed,%v",err)
	}
	// 记录交易信息
    transaction := model.Transaction{
        TransactionID:   transactionID,
        OrderID:         req.OrderId,
        UserID:          uint(req.UserId),
        Amount:          float64(req.Amount),
        CreditCardNumber: maskCreditCard(req.CreditCard.CreditCardNumber), // 仅存储后4位
        Status:          "SUCCESS",
    }
	err = mysql.CreateTransaction(&transaction)
	if err!=nil{
		return nil,fmt.Errorf("failed to save transaction, %v",err)
	}
	resp = new(payment.ChargeResp)
	resp.TransactionId = transactionID
	return 
}

// 没有第三方支付接口，这里实现一个模拟的支付接口
func processPayment(req *payment.ChargeReq)(string,error){
	// 模拟支付，如果交易金额大于1000或者卡号以4开头则交易失败
	if req.Amount > 1000 || req.CreditCard.CreditCardNumber[0] == '4' {
        return "", errors.New("payment declined by provider")
    }
	transationID := fmt.Sprintf("TX-%d", time.Now().UnixNano())
	return transationID,nil
}


// 掩码信用卡号（仅保留后4位）
func maskCreditCard(cardNumber string) string {
    if len(cardNumber) < 4 {
        return cardNumber
    }
    return "****-****-****-" + cardNumber[len(cardNumber)-4:]
}