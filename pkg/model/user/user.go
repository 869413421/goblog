package user

import "goblog/pkg/model"

type User struct {
	model.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(255) not null;unique"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}
