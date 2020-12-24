package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (user *User) Create() (err error) {
	err = model.DB.Create(&user).Error
	if err != nil {
		logger.Danger(err, "model user create error")
	}

	return
}
