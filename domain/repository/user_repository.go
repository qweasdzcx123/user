package repository

import "github.com/qweasdzcx123/user/domain/model"

type IUserRepository interface{
	//初始化数据表
	InitTable() error
	//根据用户名称查找用户信息
	FindUserByname(string)(*model.User,error)
	//根据用户ID查找用户信息
	FindUserByID(int64)(*model.User,error)
	//创建用户
	CreateUser(*model.User)(int64,error)
	//根据用户ID删除用户
	DeleteUserByID(int64) error
	//更新用户信息
	UpdateUser(*model.User)error
	//查找用户
	FindAll()([]model.User,error)
}

func NewUserRepository(db *gorm.DB) IUserRepository{
	return &UserRepository{mysqlDb:db}
}

type UserRepository struct{
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error{
	retyrb u.mysqlDb.CreateTable(&model.User{}).Error
}

//根据用户名称查找用户信息
func (u *UserRepository)FindUserByname(name string)(user *model.User,err error){
	user = &model.User{}
	return user,u.mysqlDb.where(query:"user_name = ?",name).Find(user).Error
}

	//根据用户ID查找用户信息
func (u *UserRepository)FindUserByID(userID int64)(user *model.User,err error){
		user = &model.User{}
		return user,u.mysqlDb.First(user,userID).Error
	}

//创建用户
func (u *UserRepository)CreateUser(user *model.User)(userID int64,err error){
	return u.mysqlDb.Create(user).Error
}

//根据用户ID删除用户
func (u *UserRepository)DeleteUserByID(userID int64) error{
	return u.mysqlDb.where("id = ?",userID).Delete(&model.user{}).Error
}

//更新
func (u *UserRepository)UpdateUser(user *model.User)error{
	return u.mysqlDb.Model(user).update(&user).error
}

//查找所有用户
func (u *UserRepository)FindAll()(userAll []model.User,err error){
	return userAll,u.mysqlDb.Find(&userAll).Error
}
