package article

import (
	"goblog/pkg/logger"
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

func (article *Article) Create() (err error) {
	err = model.DB.Create(&article).Error
	if err != nil {
		logger.Danger(err, "model article create error")
	}

	return
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model article update error")
		return
	}

	rowsAffected = result.RowsAffected
	return
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model article delete error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}
