package service

import (
	"eshop/dao/mysql"
	"eshop/model"
	"eshop/rpc/product/cache"
	"fmt"
	"time"
	// "gorm.io/gorm"
)

var (
	Cache *cache.RedisStruct
)

func init(){
	Cache = cache.NewRedisStruct("localhost:6379")
	// cache.
	
}

type ProductService struct{
}

func GetProduct(id uint32)(*model.Product,error){
	key:=fmt.Sprintf("product:%d",id)
	// t1:=time.Now().UnixNano()
	// if product,err:=Cache.GetProduct(key);err==nil{
	// 	fmt.Println("success cache")
	// 	fmt.Println(time.Now().UnixNano()-t1)
	// 	return product,nil
	// }

	t1 := time.Now().UnixNano()
	product,err := mysql.GetProductByID(id)
	if err!=nil{
		
		return &product,nil
	}
	
	fmt.Println(key)
	fmt.Println(time.Now().UnixNano()-t1)
	Cache.SetProduct(key,&product)
	return &product,err
}

func ListProducts(page int32,pageSize int64,categoryName string)([]model.Product,error){
	key:=fmt.Sprintf("products:%d:%d:%s",page,pageSize,categoryName)
	if products,err:=Cache.GetProducts(key);err==nil{
		return products,nil
	}
	products,err := mysql.GetProductsByCategoryWithPagination(categoryName,int(page),int(pageSize))
	if err !=nil{
		return nil,err
	}
	Cache.SetProducts(key,products)
	return products,err
}

func SearchProducts(query string,pageSize int64)([]model.Product,error){
	key:=fmt.Sprintf("search:%s%d",query,pageSize)
	if products,err:=Cache.GetProducts(key);err==nil{
		return products,nil
	}
	products,err:=mysql.SearchProducts(query,int(pageSize))
	if err!=nil{
		return nil,err
	}
	Cache.SetProducts(key,products)
	return products,nil
}