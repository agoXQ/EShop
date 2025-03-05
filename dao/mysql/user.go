package mysql

import "eshop/model"

func CreateUser(users []*model.User) error {
	return DB.Create(users).Error
}

func GetUserID(email string)(int32,error){
	db:=DB.Model(model.User{})
	user := new(model.User)
	if err:=db.Where("email = ?",email).First(&user).Error;err!=nil{
		return 0,err
	}
	return int32(user.ID),nil
}

func GetUser(email string)(*model.User,error){
	db:=DB.Model(model.User{})
	user:= new(model.User)
	if err := db.Where("email = ?",email).First(&user).Error; err != nil{
		return nil,err
	}
	return user,nil
}