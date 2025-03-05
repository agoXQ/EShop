package mysql

import (
	// "eshop/kitex_gen/product"
	"eshop/model"
)

func CreateProduct(products []*model.Product)error{
	return DB.Create(products).Error
}

// 从数据库中查询指定id的商品
 func GetProductByID(id uint32) (model.Product,error){
	var product model.Product
	err := DB.Where("id=?",id).First(&product).Error
	return product,err
 }

// 从数据库中查询指定类型的商品，支持分页查找
func GetProductsByCategoryWithPagination(categoryName string, page int, pageSize int) ([]model.Product, error) {
    var products []model.Product

    // 计算 Offset
    offset := (page - 1) * pageSize

    // 查询包含指定类别的商品
    err := DB.Where("JSON_CONTAINS(categories, ?)", `"`+categoryName+`"`).
        Offset(offset).
        Limit(pageSize).
        Find(&products).Error

    if err != nil {
        return nil, err
    }

    return products, nil
}

// SearchProducts 根据 query 搜索商品
func SearchProducts(query string,pageSize int) ([]model.Product, error) {
    var products []model.Product

    // 多字段模糊匹配
    err := DB.Where("name LIKE ? OR description LIKE ? OR JSON_CONTAINS(categories,?)", "%"+query+"%", "%"+query+"%", `"`+query+`"`).
		Limit(pageSize).
		Find(&products).Error
    if err != nil {
        return nil, err
    }

    return products, nil
}
