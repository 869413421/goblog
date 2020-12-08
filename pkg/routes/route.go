package routes

import (
	"github.com/gorilla/mux"
	"goblog/pkg/logger"
	"net/http"
)

var Router *mux.Router

func init() {
	//创建路由对象
	Router = mux.NewRouter()

	//装载routes.go中所有路由
	for _, route := range webRoutes {
		Router.Methods(route.Method).Name(route.Name).Path(route.Pattern).Handler(route.Handle)
	}
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.Danger(err, "Route Name2URL err")
		return ""
	}

	return url.String()
}

//根据名称获取路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
