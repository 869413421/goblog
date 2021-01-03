package category

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

func (category *Category) Create() (err error) {
	err = model.DB.Create(category).Error
	if err != nil {
		logger.Danger(err, "create category model error")
	}
	return err
}

func All() (categories []Category, err error) {
	err = model.DB.Find(&categories).Error
	return
}

func Get(idstr string) (category Category, err error) {
	err = model.DB.First(&category,idstr).Error
	return
}
