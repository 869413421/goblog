package routes

import (
	"goblog/app/Http/Controllers"
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
		Handle:  new(Controllers.IndexController).Home,
	},
}
