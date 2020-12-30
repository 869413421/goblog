package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func GetById(idStr string) (Article, error) {
	id := types.StringToInt(idStr)
	article := Article{}
	err := model.DB.Preload("User").First(&article, id).Error
	return article, err
}

func GetByUserID(uid string) (articles []Article, err error) {
	err = model.DB.Preload("User").Where("user_id = ?", uid).Find(&articles).Error
	return
}

func GetAll() (articles []Article, err error) {
	err = model.DB.Preload("User").Find(&articles).Error
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
