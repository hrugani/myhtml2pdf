package app

import (
	"github.com/hrugani/myhtml2pdf/webapi/app"
	"github.com/hrugani/myhtml2pdf/webapi/controller"
)

func mapUrls() {
	router.GET("/ping", controller.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
