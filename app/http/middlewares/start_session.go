package middlewares

import (
	"goblog/pkg/session"
	"net/http"
)

func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//1.启动会话
		session.StartSession(writer, request)
		//2.前往下一个请求
		next.ServeHTTP(writer, request)
	})

}
