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

var auth = new(controllers.AuthController)
var article = new(controllers.ArticlesController)
var page = new(controllers.PagesController)

var Routes = WebRoutes{
	{
		Name:    "home",
		Method:  "get",
		Pattern: "/",
		Handle:  article.Index,
	},
	{
		Name:    "about",
		Method:  "get",
		Pattern: "/about",
		Handle:  page.About,
	},
	{
		Name:    "articles.index",
		Method:  "get",
		Pattern: "/articles",
		Handle:  article.Index,
	},
	{
		Name:    "articles.show",
		Method:  "get",
		Pattern: "/articles/{id:[0-9]+}",
		Handle:  article.Show,
	},
	{
		Name:    "articles.create",
		Method:  "get",
		Pattern: "/articles/create",
		Handle:  article.Create,
	},
	{
		Name:    "articles.store",
		Method:  "POST",
		Pattern: "/articles/store",
		Handle:  article.Store,
	},
	{
		Name:    "articles.edit",
		Method:  "get",
		Pattern: "/articles/{id:[0-9]+}/edit",
		Handle:  article.Edit,
	},
	{
		Name:    "articles.update",
		Method:  "POST",
		Pattern: "/articles/{id:[0-9]+}/update",
		Handle:  article.Update,
	},
	{
		Name:    "articles.delete",
		Method:  "POST",
		Pattern: "/articles/{id:[0-9]+}/delete",
		Handle:  article.Delete,
	},
	{
		Name:    "admin.home",
		Method:  "get",
		Pattern: "/admin/home",
		Handle:  page.AdminHome,
	},
	{
		Name:    "auth.register",
		Method:  "get",
		Pattern: "/auth/register",
		Handle:  middlewares.Gust(auth.Register),
	},
	{
		Name:    "auth.do-register",
		Method:  "POST",
		Pattern: "/auth/do-register",
		Handle:  middlewares.Gust(auth.DoRegister),
	},
	{
		Name:    "auth.login",
		Method:  "get",
		Pattern: "/auth/login",
		Handle:  middlewares.Gust(auth.Login),
	},
	{
		Name:    "auth.do-login",
		Method:  "POST",
		Pattern: "/auth/do-login",
		Handle:  middlewares.Gust(auth.DoLogin),
	},
	{
		Name:    "auth.logout",
		Method:  "POST",
		Pattern: "/auth/logout",
		Handle:  middlewares.Auth(auth.Logout),
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
