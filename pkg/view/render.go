package view

import (
	"goblog/pkg/auth"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

func Render(writer io.Writer, data D, tmplFiles ...string) {
	RenderTemplate(writer, "app", data, tmplFiles...)
}

func RenderSimple(writer io.Writer, data D, tmplFiles ...string) {
	RenderTemplate(writer, "simple", data, tmplFiles...)
}

func RenderTemplate(writer io.Writer, name string, data D, tmplFiles ...string) {
	//1.获取全局数据
	data["isLogined"] = auth.Check()

	//2.生成模板解析文件
	newFiles := getTemplateFiles(tmplFiles...)

	//3.挂载所有模板
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	if err != nil {
		logger.Danger(err, "render error")
	}
	//渲染模板
	tmpl.ExecuteTemplate(writer, name, data)

}

func getTemplateFiles(templateFiles ...string) []string {
	//1.定义视图根目录
	viewDir := "resources/views/"

	//2.将需要渲染的模板修改为相对路径
	for i, name := range templateFiles {
		name = strings.Replace(name, ".", "/", -1)
		templateFiles[i] = viewDir + name + ".html"
	}

	//3.获取所有布局模板的路径
	files, err := filepath.Glob(viewDir + "layouts/*.html")

	if err != nil {
		logger.Danger(err, "render error")
	}

	newFiles := append(files, templateFiles...)

	return newFiles
}
