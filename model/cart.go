package model

import "gorm.io/gorm"

type Cart struct{
	gorm.Model
	UserID uint
	ProductID uint
	Quantity int
}

func (c *Cart) TableName()string{
	return "carts"
}