package Controllers

import (
	"fmt"
	"net/http"
)

type IndexController struct {
}

func (controller *IndexController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
