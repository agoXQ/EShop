package model

import "gorm.io/gorm"


type Order struct {
    gorm.Model
    OrderID       string      `gorm:"primaryKey;column:order_id;type:varchar(36);not null"` // 订单ID
    UserID        uint        `gorm:"column:user_id;not null"`                              // 用户ID
    UserCurrency  string      `gorm:"column:user_currency;type:varchar(10);not null"`       // 用户使用的货币类型
    StreetAddress 	  string      `gorm:"column:street_address;type:varchar(255);not null"`     // 收货地址
    City          string      `gorm:"column:city;type:varchar(100);not null"`               // 城市
    State         string      `gorm:"column:state;type:varchar(100);not null"`              // 州/省份
    Country       string      `gorm:"column:country;type:varchar(100);not null"`            // 国家
    ZipCode       int         `gorm:"column:zip_code;not null"`                             // 邮政编码
    Email         string      `gorm:"column:email;type:varchar(255);not null"`              // 用户邮箱
    CreatedAt     int64       `gorm:"column:created_at;not null"`                           // 订单创建时间（时间戳）
    OrderItems    []OrderItem `gorm:"foreignKey:OrderID;references:OrderID"`                // 订单项列表
}

func (o *Order) TableName() string {
    return "orders"
}


type OrderItem struct {
    gorm.Model
    ID           uint    `gorm:"primaryKey;autoIncrement"`                     // 订单项ID
    OrderID      uint  `gorm:"column:order_id;type:varchar(36);not null"`     // 订单ID（外键）
    ItemID       uint32  `gorm:"column:item_id;type:varchar(36);not null"`      // 商品ID
    ItemName     string  `gorm:"column:item_name;type:varchar(255);not null"`   // 商品名称
    ItemQuantity int32    `gorm:"column:item_quantity;not null"`                 // 商品数量
    ItemCost     float32 `gorm:"column:item_cost;type:decimal(10,2);not null"`  // 商品价格
}

func (oi *OrderItem) TableName() string {
    return "order_items"
}

