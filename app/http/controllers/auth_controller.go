package controllers

import (
	"fmt"
	"goblog/app/http/requests"
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
		fmt.Fprint(w, "创建用户失败")
		return
	}

	fmt.Fprint(w, "创建用户成功，ID:"+_user.GetStringID())
}
