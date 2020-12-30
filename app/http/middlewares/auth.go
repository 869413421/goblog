package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !auth.Check() {
			flash.Warning("请登陆后再进行访问")
			http.Redirect(writer, request, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(writer, request)
	}
}
