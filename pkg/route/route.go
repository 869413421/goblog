package route

import (
	"github.com/gorilla/mux"
	"goblog/pkg/logger"
	"net/http"
)

var route *mux.Router

// SetRoute 设置路由实例，以供 Name2URL 等函数使用
func SetRoute(r *mux.Router) {
	route = r
}

//根据名称获取路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {

	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		logger.Danger(err, "name to url error")
		return ""
	}
	return  url.String()
}
