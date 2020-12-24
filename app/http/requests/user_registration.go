package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/pkg/model/user"
)

func ValidateRegistrationForm(data user.User) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{
			"required",
			"alpha_num",
			"between:3,30",
			"not_exists:users,name",
		},
		"email": []string{
			"required",
			"email",
			"between:3,30",
			"not_exists:users,email",
		},
		"password": []string{
			"required",
			"between:6,30",
		},
		"password_comfirm": []string{
			"required",
			"between:6,30",
		},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required：用户名为必填选项",
			"alpha_num:只允许数字和英文",
			"between:用户名在3到30个字符之间",
		},
		"email": []string{
			"required:Email必填",
			"email:邮件格式错误",
			"between:邮件长度在3到30字符之间",
		},
		"password": []string{
			"required:密码必填",
			"between:密码在3到30个字符之间",
		},
		"password_comfirm": []string{
			"required:确认密码框为必填项",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.PasswordComfirm {
		errs["password_comfirm"] = append(errs["password_comfirm"], "两次输入的密码不一致")
	}

	return errs
}
