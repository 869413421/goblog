package article

import (
	"goblog/pkg/model"
	"goblog/pkg/route"
)

type Article struct {
	model.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", article.GetStringID())
}
