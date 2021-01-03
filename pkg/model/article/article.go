package article

import (
	"goblog/pkg/model"
	"goblog/pkg/model/category"
	"goblog/pkg/model/user"
	"goblog/pkg/route"
)

type Article struct {
	model.BaseModel
	Title      string `gorm:"column:title;type:varchar(255);not null" valid:"title"`
	Body       string `gorm:"column:body;type:text;not null" valid:"body"`
	UserID     uint64 `gorm:"not null;index`
	User       user.User
	CategoryID uint64 `gorm:"not null;default:1;index"`
	Category   category.Category
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", article.GetStringID())
}
