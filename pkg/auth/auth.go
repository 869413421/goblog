package auth

import (
	"errors"
	"fmt"
	"goblog/pkg/model/user"
	"goblog/pkg/session"
	"gorm.io/gorm"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if len(uid) > 0 && ok {
		return uid
	}

	return ""
}

//获取登陆用户
func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		user, err := user.GetByID(uid)
		if err == nil {
			return user
		}
	}

	return user.User{}
}

func Attempt(email string, password string) error {
	//1.根据邮箱获取用户
	user, err := user.GetByEmail(email)
	fmt.Println("==========读取数据库===========")
	fmt.Println(err)
	//2.判断获取用户信息是否出错
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账户不存在")
		} else {
			return errors.New("内部错误")
		}
	}

	//3.比较用户输入密码
	if !user.ComparePassword(password) {
		return errors.New("账户不存在,或者密码错误")
	}

	Login(user)

	return nil
}

//登陆用户
func Login(user user.User) {
	session.Put("uid", user.GetStringID())
}

//退出
func Logout() {
	session.Forget("uid")
}

//检查是否登陆
func Check() bool {
	return len(_getUID()) > 0
}
