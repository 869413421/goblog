package routes

import (
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

var webRoutes = WebRoutes{
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
}

var notFountHandler = http.HandlerFunc(new(controllers.PagesController).NotFound)
