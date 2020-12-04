package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	//创建路由对象
	router := mux.NewRouter()

	//装载routes.go中所有路由
	for _, route := range webRoutes {
		router.Methods(route.Method).Name(route.Name).Path(route.Pattern).Handler(route.Handle)
	}
	//返回对象指针
	return router
}
