package main

import (
	"context"
	"encoding/json"
	// "eshop/dao/mysql"
	product "eshop/kitex_gen/product"
	"eshop/model"
	"eshop/rpc/product/service"
	"fmt"
	// "log"
	"strings"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// TODO: Your code here...
	// category := req.CategoryName
	// page :=req.Page 
	// pageSize := req.PageSize
	// products,err :=mysql.GetProductsByCategoryWithPagination(category,int(page),int(pageSize))
	// if err!=nil{
	// 	return nil,err
	// }
	// resp = new(product.ListProductsResp)
	
	// resp.Products = ConvertModelProductsToProtoProducts(products)
	// log.Printf("%v",products)

	pros,err := service.ListProducts(req.Page,req.PageSize,req.CategoryName)
	if err != nil{
		return nil,fmt.Errorf("fail to list product,%s",req.CategoryName)
	}
	var protoProducts []*product.Product
	for _, p := range pros{
		protoProducts = append(protoProducts,&product.Product{
			Id: uint32(p.ID),
			Name: p.Name,
			Description: p.Description,
			Price: p.Price,
			Picture: p.Picture,
			Categories: splitCategories(p.Categories),
		})
	}
	resp = new(product.ListProductsResp)
	resp.Products = protoProducts
	return
}

func ConvertModelProductToProtoProduct(modelProduct model.Product) *product.Product {
	var res = &product.Product{
        Id:          uint32(modelProduct.ID), // 注意：modelProduct.ID 是 uint 类型，需要转换为 uint32
        Name:        modelProduct.Name,
        Description: modelProduct.Description,
        Picture:     modelProduct.Picture,
        Price:       modelProduct.Price,
    }
	json.Unmarshal([]byte(modelProduct.Categories), &res.Categories)
	return res
}

func ConvertModelProductsToProtoProducts(modelProducts []model.Product) []*product.Product {
    protoProducts := make([]*product.Product, len(modelProducts))
    for i, modelProduct := range modelProducts {
        protoProducts[i] = ConvertModelProductToProtoProduct(modelProduct)
    }
    return protoProducts
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// TODO: Your code here...
	// ID := req.Id
	// p,err := mysql.GetProductByID(ID)
	// resp = new(product.GetProductResp)
	// resp.Product = ConvertModelProductToProtoProduct(p)
	pro,err :=service.GetProduct(req.Id)
	if err !=nil{
		return nil,fmt.Errorf("fail to get product : %v",req.Id)
	}

	resp = new(product.GetProductResp)
	resp.Product = ConvertModelProductToProtoProduct(*pro)
	return
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// TODO: Your code here...
	// resp = new(product.SearchProductsResp)
	// // req.Query
	// products,err := mysql.SearchProducts(req.Query,defaultPageSize)
	// resp.Results = ConvertModelProductsToProtoProducts(products)
	pros,err := service.SearchProducts(req.Query,50)
	if err!=nil{
		return nil,fmt.Errorf("failed to search product, %s",req.Query)
	}
	var protoProducts []*product.Product
	for _, p:=range pros{
		protoProducts = append(protoProducts, &product.Product{
			Id:uint32(p.ID),
			Name:p.Name,
			Description: p.Description,
			Picture: p.Picture,
			Price: p.Price,
			Categories: splitCategories(p.Categories),
		})
	}
	resp = new(product.SearchProductsResp)
	resp.Results = protoProducts
	return
}


func splitCategories(categories string) []string {
    // 将逗号分隔的字符串转换为切片
    return strings.Split(categories, ",")
}