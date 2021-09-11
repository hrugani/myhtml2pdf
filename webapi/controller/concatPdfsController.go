package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hrugani/myhtml2pdf/services"
)

func Concat(c *gin.Context) {

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		msgErr := fmt.Sprintf("[ERROR] receiving request. detail: %s ", err.Error())
		log.Default().Print(msgErr)
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": msgErr})
		return
	}
	files := form.File["files"]
	log.Default().Println("[INFO] multipart form request was received successfully")

	err = validatePayload(files)
	if err != nil {
		msgErr := "[ERROR] validating request payload. detail: " + err.Error()
		log.Default().Println(msgErr)
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": msgErr})
		return
	}
	log.Default().Println("[INFO]", "request payload was validated")

	// creates workdir
	workDirName, err := createWorkDir()
	if err != nil {
		msgErr := "[ERROR] creating workdir. detail: " + err.Error()
		log.Default().Println(msgErr)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": msgErr})
		return
	}
	log.Default().Printf("[INFO] workdir %s was created", workDirName)

	// save uploaded file in the workdir
	var uploadedFileName string
	for _, file := range files {
		pathSep := fmt.Sprintf("%c", os.PathSeparator)
		uploadedFileName = filepath.Base(file.Filename)
		fileNameInWorkDir := workDirName + pathSep + uploadedFileName
		if err := c.SaveUploadedFile(file, fileNameInWorkDir); err != nil {
			msgErr := fmt.Sprintf("[ERROR] saving uploaded file %s in workdir. detail: %s", fileNameInWorkDir, err.Error())
			log.Default().Printf(msgErr)
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": msgErr})
			return
		}
	}
	log.Default().Printf("[INFO] file %v was uploaded and saved properly in the server in the workdir", uploadedFileName)

	// process the html to pdf convertion
	var pdfFilePath string
	pdfFilePath, err = services.ConcatPdfs(workDirName, uploadedFileName)
	if err != nil {
		msgErr := fmt.Sprintf("[ERROR] executing concatpdfs service, detail: %s ", err.Error())
		log.Default().Println(msgErr)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": msgErr})
		return
	}
	log.Default().Print("[INFO]", "concatPdfs service was executed")

	// reponse: pdf file
	c.File(string(pdfFilePath))
	log.Default().Printf("[INFO] concat pdf executed successfully. pdf file generated: %s", pdfFilePath)

	// Removes workdir
	err = removeWorkDir(workDirName)
	if err != nil {
		log.Default().Printf("[Error] deliting used workdir %s. detail: %s", workDirName, err.Error())
	}
	log.Default().Printf("[INFO] executed clean up of workdir: %s", workDirName)
}
