package app

import (
	"github.com/hrugani/myhtml2pdf/webapi/controller"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)
	router.GET("/version", controller.Version)
	router.POST("/html2pdf", controller.ConvertHtml2PDF)
	router.POST("/merge", controller.MergePDFs)
}
