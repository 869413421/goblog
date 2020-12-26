package controllers

import (
	"goblog/app/http/requests"
	"goblog/pkg/auth"
	"goblog/pkg/model/user"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (controller *AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//1.验证数据
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	password_comfirm := r.PostFormValue("password_comfirm")

	_user := user.User{
		Name:            name,
		Email:           email,
		Password:        password,
		PasswordComfirm: password_comfirm,
	}

	//2.失败跳转回注册视图
	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
		return
	}

	//3.验证成功入库，跳转首页
	_user.Create()
	if _user.ID <= 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
		return
	}

	auth.Login(_user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

func (controller *AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	//1.获取表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	//2.用户认证
	err := auth.Attempt(email, password)

	//3.根据认证成功状态跳转
	if err != nil {
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
		return
	}

	http.Redirect(w, r, "/articles", http.StatusFound)
}

func (controller *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	http.Redirect(w, r, "/", http.StatusFound)
}
