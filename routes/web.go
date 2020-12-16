package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

type WebRouter struct {
	Name    string
	Method  string
	Pattern string
	Handle  http.HandlerFunc
}

type WebRoutes []WebRouter

var Routes = WebRoutes{
	{
		Name:    "home",
		Method:  "get",
		Pattern: "/",
		Handle:  new(controllers.PagesController).Home,
	},
	{
		Name:    "about",
		Method:  "get",
		Pattern: "/about",
		Handle:  new(controllers.PagesController).About,
	},
	{
		Name:    "articles.index",
		Method:  "get",
		Pattern: "/articles",
		Handle:  new(controllers.ArticlesController).Index,
	},
	{
		Name:    "articles.show",
		Method:  "get",
		Pattern: "/articles/{id:[0-9]+}",
		Handle:  new(controllers.ArticlesController).Show,
	},
	{
		Name:    "articles.create",
		Method:  "get",
		Pattern: "/articles/create",
		Handle:  new(controllers.ArticlesController).Create,
	},
	{
		Name:    "admin.home",
		Method:  "get",
		Pattern: "/admin/home",
		Handle:  new(controllers.PagesController).AdminHome,
	},
}

var NotFountHandler = http.HandlerFunc(new(controllers.PagesController).NotFound)

func RegisterWebRoutes(router *mux.Router) {
	//装载所有路由
	for _, route := range Routes {
		router.Methods(route.Method).Name(route.Name).Path(route.Pattern).Handler(route.Handle)
	}
}
