package app

import (
	"github.com/hrugani/myhtml2pdf/webapi/controller"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)
	router.POST("/convert", controller.Convert)
}
