package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/pkg/model/category"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
