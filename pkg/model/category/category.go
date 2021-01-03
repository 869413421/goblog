package category

import (
	"goblog/pkg/model"
	"goblog/pkg/route"
)

type Category struct {
	model.BaseModel
	Name string `gorm:"type:varchar(255);not null" valid:"name"`
}

// Link 方法用来生成文章链接
func (category Category) Link() string {
	return route.Name2URL("categories.show", "id", category.GetStringID())
}
