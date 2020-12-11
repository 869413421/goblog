package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
	"sync"
)

var Router *mux.Router
var once sync.Once

func SetupRoute() *mux.Router {
	once.Do(func() {
		Router = mux.NewRouter()
		routes.RegisterWebRoutes(Router)
	})
	return Router
}
