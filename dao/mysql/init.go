package mysql

import (
	"eshop/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var dsn = "gorm:gorm@tcp(localhost:9911)/gorm?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

// func Init(){
// 	var err error
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 		PrepareStmt:            true,
// 		Logger:                 logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// }

func InitDB() {
    // dsn := "user:password@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // 自动迁移表结构
    if err := db.AutoMigrate(&model.User{},&model.Product{},&model.Cart{},&model.Order{},&model.Transaction{}); err != nil {
        log.Fatalf("Failed to migrate tables: %v", err)
    }

    log.Println("Database connected and tables migrated successfully")
    DB = db
}