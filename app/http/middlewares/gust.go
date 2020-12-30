package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

func Gust(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if auth.Check() {
			flash.Warning("只允许游客访问")
			http.Redirect(writer, request, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(writer, request)
	}
}
