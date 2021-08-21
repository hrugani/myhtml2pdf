package controller

import (
	"fmt"
	"net/http"
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

		fmt.Printf("file received:  %#v \n\n", file)
		c.String(
			http.StatusOK,
			fmt.Sprintf("file received:  %#v \n\n", file),
		)

		filename := filepath.Base(file.Filename + "_test111111")
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
