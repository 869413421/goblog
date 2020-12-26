package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
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
		Handle:  new(controllers.ArticlesController).Index,
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
		Name:    "articles.store",
		Method:  "POST",
		Pattern: "/articles/store",
		Handle:  new(controllers.ArticlesController).Store,
	},
	{
		Name:    "articles.edit",
		Method:  "get",
		Pattern: "/articles/{id:[0-9]+}/edit",
		Handle:  new(controllers.ArticlesController).Edit,
	},
	{
		Name:    "articles.update",
		Method:  "POST",
		Pattern: "/articles/{id:[0-9]+}/update",
		Handle:  new(controllers.ArticlesController).Update,
	},
	{
		Name:    "articles.delete",
		Method:  "POST",
		Pattern: "/articles/{id:[0-9]+}/delete",
		Handle:  new(controllers.ArticlesController).Delete,
	},
	{
		Name:    "admin.home",
		Method:  "get",
		Pattern: "/admin/home",
		Handle:  new(controllers.PagesController).AdminHome,
	},
	{
		Name:    "auth.register",
		Method:  "get",
		Pattern: "/auth/register",
		Handle:  new(controllers.AuthController).Register,
	},
	{
		Name:    "auth.do-register",
		Method:  "POST",
		Pattern: "/auth/do-register",
		Handle:  new(controllers.AuthController).DoRegister,
	},
	{
		Name:    "auth.login",
		Method:  "get",
		Pattern: "/auth/login",
		Handle:  new(controllers.AuthController).Login,
	},
	{
		Name:    "auth.do-login",
		Method:  "POST",
		Pattern: "/auth/do-login",
		Handle:  new(controllers.AuthController).DoLogin,
	},
	{
		Name:    "auth.logout",
		Method:  "POST",
		Pattern: "/auth/logout",
		Handle:  new(controllers.AuthController).Logout,
	},
}

var NotFountHandler = http.HandlerFunc(new(controllers.PagesController).NotFound)

func RegisterWebRoutes(router *mux.Router) {
	//装载所有路由
	for _, route := range Routes {
		router.Methods(route.Method).Name(route.Name).Path(route.Pattern).Handler(route.Handle)
	}
	//全局中间件
	//router.Use(middlewares.ForceHTML)
	router.Use(middlewares.StartSession)
}
