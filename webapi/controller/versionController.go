package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "mypdfservices - current version - v0.5.2"})
}
