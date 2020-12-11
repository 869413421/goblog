package controllers

import (
	"fmt"
	"goblog/pkg/view"
	"html/template"
	"net/http"
)

type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	view.GenerateHTML(w, nil, "layout", "home")
}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

// 后台首页
func (*PagesController) AdminHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("resources/views/layout.admin.html")
	t.Execute(w, nil)
}
