package controllers

import (
	"fmt"
	"goblog/pkg/logger"
	"goblog/pkg/model/article"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"unicode/utf8"
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

type ArticlesFormData struct {
	Title, Body string
	URL         string
	Errors      map[string]string
}

func (controller *ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeUrl := route.Name2URL("articles.store")
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeUrl,
		Errors: nil,
	}
	view.GenerateHTML(w, data, "layout.admin", "article/edit")
}
func ValidateArticlesFromData(title string, body string) map[string]string {
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "标题不允许为空"
	} else if utf8.RuneCountInString(title) < 2 || utf8.RuneCountInString(title) > 100 {
		errors["title"] = "标题长度需介于 3-40"
	}

	if body == "" {
		errors["body"] = "内容不允许为空"
	} else if utf8.RuneCountInString(body) < 3 || utf8.RuneCountInString(body) > 2000 {
		errors["body"] = "内容长度需介于 3-2000"
	}
	return errors
}
func (controller *ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := ValidateArticlesFromData(title, body)
	if len(errors) > 0 {
		storeUrl := route.Name2URL("articles.create")
		data := ArticlesFormData{
			Title:  "",
			Body:   "",
			URL:    storeUrl,
			Errors: errors,
		}

		view.GenerateHTML(w, data, "layout.admin", "article/edit")
	}

	_article := article.Article{
		Title: title,
		Body:  body,
	}

	err := _article.Create()
	if err != nil || _article.ID <= 0 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "创建失败，服务器错误")
	}
	fmt.Fprintf(w, "创建成功，ID:"+strconv.FormatInt(_article.ID, 10))
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
	deleteUrl := route.Name2URL("articles.delete", "id", id)
	data := ArticlesFormData{
		Title:  _article.Title,
		Body:   _article.Body,
		URL:    deleteUrl,
		Errors: nil,
	}
	view.GenerateHTML(w, data, "layout.admin", "article/show")
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
	updateUrl := route.Name2URL("articles.update", "id", id)
	data := ArticlesFormData{
		Title:  _article.Title,
		Body:   _article.Body,
		URL:    updateUrl,
		Errors: nil,
	}

	view.GenerateHTML(w, data, "layout.admin", "article/edit")
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

	errors := ValidateArticlesFromData(title, body)
	if len(errors) > 0 {
		updateUrl := route.Name2URL("articles.edit", "id", id)
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    updateUrl,
			Errors: errors,
		}

		view.GenerateHTML(w, data, "layout.admin", "article/edit")
	}
	_article.Title = title
	_article.Body = body
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
