package user

import "goblog/pkg/model"

type User struct {
	model.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255) not null;unique" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255);not null" valid:"password"`
	// gorm:"-" 使用这个注解GORM读写会忽略这个字段
	PasswordComfirm string `gorm:"-" valid:"password_comfirm"`
}
