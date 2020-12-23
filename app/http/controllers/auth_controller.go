package controllers

import (
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "auth.register")
}

func (controller *AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//1.验证数据

	//2.失败跳转回注册视图

	//3.验证成功入库，跳转首页
}
