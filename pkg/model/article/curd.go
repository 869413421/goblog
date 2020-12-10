package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func GetById(idStr string) (article Article, err error) {
	id := types.StringToInt(idStr)
	err = model.DB.First(&article, id).Error
	return
}

func GetAll() (articles []Article, err error) {
	err = model.DB.Find(&articles).Error
	return
}
