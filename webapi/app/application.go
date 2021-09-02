package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication(port int) {
	listeningString := fmt.Sprintf(":%d", port)
	mapUrls()
	router.Run(listeningString)
}
