package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name      string `json:"name" column:"name"`
	Description    string  `json:"description" column:"description"`
	Picture string `json:"picture" column:"picture"`
	Price float32 	`json:"price" column:"price"`
	Categories string `json:"categories" column:"cotegories" gorm:"type:json"`
}

func (p *Product) TableName() string {
	return "products"
}