package model
type User struct{
	//主键
	ID int64 'gorm:"primary;not_null;auto_increment"'
	//用户名称
	UserName string 'gorm:"unique_idex;not_null"'
	//需要添加的字段
	FirstName string
	//....
	//密码
	HashPassword sring
}