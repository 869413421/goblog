package controllers

import (
	"fmt"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

type BaseController struct {
}

func (controller BaseController) ResponseForSqlError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "数据未找到")
		return
	} else {
		logger.Danger(err, "controller error")
	}
}

func (controller BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("未授权操作")
	http.Redirect(w, r, "/", http.StatusFound)
}
