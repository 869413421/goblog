package controllers

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model/article"
	"goblog/pkg/model/user"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
}

func (controller UserController) Show(w http.ResponseWriter, r *http.Request) {
	//1.获取URL参数
	id := route.GetRouteVariable("id", r)

	//2.查找用户
	_user, err := user.GetByID(id)

	//3.如果发生错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//3.1 找不到用户
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "找不到用户")
			return
		} else {
			//3.2 内部错误
			logger.Danger(err, "user show controller error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "内部错误")
			return
		}
	}

	//4 显示用户文章
	articles, err := article.GetByUserID(_user.GetStringID())

	//4.1 如果发生错误
	if err != nil {
		logger.Danger(err, "user show controller error")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "内部错误")
		return
	}

	view.Render(w, view.D{
		"Articles": articles,
	}, "article.index", "article._article_meta")
}
