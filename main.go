package main

import (
	"fmt"
	. "goblog/config"
	"goblog/pkg/routes"
	"net/http"
	"strings"
)

func defaultHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 goBlog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func forceMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json;charset=utf-8")
		writer.Header().Set("Token", "123456")
		handler.ServeHTTP(writer, request)
	})
}

func trimUrlPath(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			request.URL.Path = strings.TrimSuffix(request.URL.Path, "/")
		}
		handler.ServeHTTP(writer, request)
	})
}

func main() {
	startWebServer()
}

func startWebServer() {
	fmt.Println("Server Start")
	//初始化配置
	config := LoadConfig()
	//初始化路由
	router := routes.Router
	//处理静态资源
	assets := http.FileServer(http.Dir(config.App.Static))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))
	//监听端口
	err := http.ListenAndServe(config.App.Address, router)
	fmt.Println("Server Running")
	if err != nil {
		fmt.Println("Start Service Error ", err)
	}
}
