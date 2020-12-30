package policies

import (
	"goblog/pkg/auth"
	"goblog/pkg/model/article"
)

//是否拥有文章权限
func CanModifyArticle(article article.Article) bool {
	return auth.User().ID == article.UserID
}
