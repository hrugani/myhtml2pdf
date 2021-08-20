package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Convert(c *gin.Context) {
	c.String(http.StatusOK, "Convert was called...")
}
