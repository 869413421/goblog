package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
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

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
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

// GetByCategoryID 获取分类相关的文章
func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("categories.show", "id", cid), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}
