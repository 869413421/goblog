package controllers

import (
	"goblog/pkg/logger"
	"net/http"
)

type IndexController struct {
}

func (controller *IndexController) Home(w http.ResponseWriter, r *http.Request) {
	logger.Info("logs")
}
