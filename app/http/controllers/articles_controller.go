package controllers

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model/article"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type ArticlesController struct {
}

func (controller *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	data, err := article.GetAll()
	if err != nil {
		logger.Danger(err, "ArticlesController Index Error")
		fmt.Fprintln(w, err)
	}
	view.GenerateHTML(w, data, "layout.admin", "article/index")
}

func (controller *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.GenerateHTML(w, nil, "layout.admin", "article/edit")
}

func (controller *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	art, err := article.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//未找到文章
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "文章未找到")
		} else
		{
			logger.Danger(err, "ArticlesController Show Error")
			fmt.Fprintln(w, err)
		}
	}
	fmt.Fprint(w, art)
}
