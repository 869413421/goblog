package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func GetByID(idStr string) (user User, err error) {
	id := types.StringToInt(idStr)
	err = model.DB.First(&user, id).Error
	return
}

func GetByEmail(emil string) (user User, err error) {
	err = model.DB.Where("email = ?", emil).First(&user).Error
	return
}

func (user *User) Create() (err error) {
	err = model.DB.Create(&user).Error
	if err != nil {
		logger.Danger(err, "model user create error")
	}

	return
}
