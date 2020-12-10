package controllers

import (
	"fmt"
	"goblog/pkg/model/article"
	"goblog/pkg/routes"
	"gorm.io/gorm"
	"net/http"
)

type ArticlesController struct {
}

func (controller *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {

}

func (controller *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := routes.GetRouteVariable("id", r)
	art, err := article.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//未找到文章
		}
	}
	fmt.Fprint(w, art)
}
