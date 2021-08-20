package controller

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Convert(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(
			http.StatusBadRequest,
			fmt.Sprintf("request err: %s", err.Error()+"\n"),
		)
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(
				http.StatusBadRequest,
				fmt.Sprintf("upload file err: %s", err.Error()+"\n"),
			)
			return
		}
	}

	c.String(
		http.StatusOK,
		fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s\n", len(files),
			name,
			email))
}
