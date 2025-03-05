package mysql

import "eshop/model"

func CreateCart(cart *model.Cart)error{
	return DB.Create(cart).Error
}

func GetCartByID(userID uint32)(*[]model.Cart,error){
	res := new([]model.Cart)
	err:=DB.Where("userid = ?",userID).Find(res).Error
	return res,err
}

func DeleteCartByID(userID uint32)error{
	return DB.Where("userid = ?",userID).Delete(&model.Cart{}).Error
}

