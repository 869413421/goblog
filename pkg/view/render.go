package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(writer io.Writer, data interface{}, tmplFiles ...string) {
	//1.定义视图根目录
	viewDir := "resources/views/"

	//2.将需要渲染的模板修改为相对路径
	for i, name := range tmplFiles {
		name = strings.Replace(name, ".", "/", -1)
		tmplFiles[i] = viewDir + name + ".html"
	}

	//3.获取所有布局模板的路径
	files, err := filepath.Glob(viewDir + "layouts/*.html")

	if err != nil {
		logger.Danger(err, "render error")
	}

	//4.挂载所有模板，和需要的模板方法
	newFiles := append(files, tmplFiles...)
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	if err != nil {
		logger.Danger(err, "render error")
	}
	//5.渲染模板
	tmpl.ExecuteTemplate(writer, "app", data)

}
