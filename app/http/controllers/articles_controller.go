package controllers

import (
	"fmt"
	"goblog/app/http/requests"
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
	articles, err := article.GetAll()
	if err != nil {
		logger.Danger(err, "ArticlesController Index Error")
		fmt.Fprintln(w, err)
	}

	view.Render(w, view.D{
		"Articles": articles,
	}, "article.index", "article._article_meta")
}

func (controller *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "article.create", "article._form_field")
}

func (controller *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	_article := article.Article{
		Title: title,
		Body:  body,
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

	view.Render(w, view.D{
		"Article": _article,
	}, "article.show", "article._article_meta")
}

func (controller *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.GetById(id)
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

	// 4. 读取成功，显示编辑文章表单
	view.Render(w, view.D{
		"Article": _article,
		"Errors":  view.D{},
	}, "article.edit", "article._form_field")

}

func (controller *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.GetById(id)
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
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "找不到文章")
		}
		w.WriteHeader(http.StatusInternalServerError)
		logger.Danger(err, "articlesController error")
		fmt.Fprint(w, "服务器错误")
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
