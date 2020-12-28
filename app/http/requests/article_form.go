package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/pkg/model/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{
			"required",
			"between:2,50",
		},
		"body": []string{
			"required",
			"min:10",
		},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
			"min:标题长度需大于 3",
			"max:标题长度需小于 40",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min:长度需大于 10",
		},
	}

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
