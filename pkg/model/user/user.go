package user

import "goblog/pkg/model"

type User struct {
	model.BaseModel
	Name     string
	Email    string
	Password string
}
