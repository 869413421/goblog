package main

import (
	"fmt"
	"goblog/bootstrap"
	. "goblog/config"
	"net/http"
)

func main() {
	startWebServer()
}

func startWebServer() {
	fmt.Println("Server Start")
	//初始化配置
	config := LoadConfig()
	//初始化路由
	router := bootstrap.SetupRoute()
	//处理静态资源
	assets := http.FileServer(http.Dir(config.App.Static))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))
	//初始化gorm
	bootstrap.SetupDB()
	//监听端口
	err := http.ListenAndServe(config.App.Address, router)
	fmt.Println("Server Running")
	if err != nil {
		fmt.Println("Start Service Error ", err)
	}
}
