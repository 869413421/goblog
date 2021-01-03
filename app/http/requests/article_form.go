package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/pkg/model/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"title": []string{"required", "min_cn:3", "max_cn:40"},
		"body":  []string{"required", "min_cn:10"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min_cn:长度需大于 10",
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
