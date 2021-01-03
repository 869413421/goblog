package controllers

import (
	"fmt"
	"goblog/app/http/requests"
	"goblog/app/policies"
	"goblog/pkg/auth"
	"goblog/pkg/model/article"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type ArticlesController struct {
	BaseController
}

func (controller *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// 1. 获取结果集
	articles, pagerData, err := article.GetAll(r, 2)

	if err != nil {
		controller.ResponseForSqlError(w, err)
	} else {

		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "article.index", "article._article_meta")
	}
}

func (controller *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "article.create", "article._form_field")
}

func (controller *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	_article := article.Article{
		Title:  title,
		Body:   body,
		UserID: auth.User().ID,
	}
	errors := requests.ValidateArticleForm(_article)
	if len(errors) > 0 {
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  errors,
		}, "article.create", "article._form_field")
	}

	err := _article.Create()
	if err != nil || _article.ID <= 0 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "创建失败，服务器错误")
	}
	http.Redirect(w, r, route.Name2URL("articles.show", "id", _article.GetStringID()), http.StatusFound)
}

func (controller *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.GetById(id)
	if err != nil {
		controller.ResponseForSqlError(w, err)
	}

	view.Render(w, view.D{
		"Article":          _article,
		"CanModifyArticle": policies.CanModifyArticle(_article),
	}, "article.show", "article._article_meta")
}

func (controller *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.GetById(id)
	if err != nil {
		controller.ResponseForSqlError(w, err)
	}

	if !policies.CanModifyArticle(_article) {
		controller.ResponseForUnauthorized(w, r)
	}

	view.Render(w, view.D{
		"Article": _article,
		"Errors":  view.D{},
	}, "article.edit", "article._form_field")

}

func (controller *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.GetById(id)
	if err != nil {
		controller.ResponseForSqlError(w, err)
	}

	if !policies.CanModifyArticle(_article) {
		controller.ResponseForUnauthorized(w, r)
	}

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	_article.Title = title
	_article.Body = body

	errors := requests.ValidateArticleForm(_article)
	if len(errors) > 0 {
		// 4.3 表单验证不通过，显示理由
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  errors,
		}, "article.edit", "article._form_field")
	}

	rowsAffected, err := _article.Update()

	if err != nil {
		// 数据库错误
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
		return
	}

	// √ 更新成功，跳转到文章详情页
	if rowsAffected > 0 {
		showURL := route.Name2URL("articles.show", "id", id)
		http.Redirect(w, r, showURL, http.StatusFound)
	} else {
		fmt.Fprint(w, "您没有做任何更改！")
	}
}

func (controller *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.GetById(id)
	if err != nil {
		controller.ResponseForSqlError(w, err)
	}

	if !policies.CanModifyArticle(_article) {
		controller.ResponseForUnauthorized(w, r)
	}

	rowsAffected, err := _article.Delete()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "服务器错误")
	}
	if rowsAffected <= 0 {
		fmt.Fprint(w, "删除失败")
	}

	indexUrl := route.Name2URL("articles.index")
	http.Redirect(w, r, indexUrl, http.StatusFound)
}
