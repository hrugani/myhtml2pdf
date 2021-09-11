package controller

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hrugani/myhtml2pdf/services"
	uuid "github.com/satori/go.uuid"
)

func Convert(c *gin.Context) {

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
		log.Default().Print(msgErr)
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": msgErr})
		return
	}
	log.Default().Println("[INFO]", "request payload was validated")

	preffix := c.Query(workdirPreffixParamName)

	// creates workdir
	workDirName, err := createWorkDir("html2pdf", preffix)
	if err != nil {
		msgErr := "[ERROR] creating workdir. detail: " + err.Error()
		log.Default().Print(msgErr)
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
			log.Default().Print(msgErr)
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": msgErr})
			return
		}
	}
	log.Default().Printf("[INFO] file %v was uploaded and saved properly in the server in the workdir", uploadedFileName)

	// process the html to pdf convertion
	var pdfFilePath string
	if strings.HasSuffix(strings.ToLower(uploadedFileName), ".zip") {
		pdfFilePath, err = services.Zip2Pdf(workDirName, uploadedFileName)
		if err != nil {
			msgErr := fmt.Sprintf("[ERROR] executing zip2pdf service, detail: %s ", err.Error())
			log.Default().Print(msgErr)
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": msgErr})
			return
		}
	} else {
		errMsg := "[ERROR] file uploaded must be .zip"
		log.Default().Print(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusInternalServerError, "message": errMsg})
		return
	}
	log.Default().Println("[INFO]", "zip to pdf service was executed")

	// reponse: pdf file
	c.File(string(pdfFilePath))
	log.Default().Printf("[INFO] html to pdf executed successfully. pdf file generated: %s", pdfFilePath)

	// Removes workdir
	err = removeWorkDir(workDirName)
	if err != nil {
		log.Default().Printf("[Error] deliting used workdir %s. detail: %s", workDirName, err.Error())
	}
	log.Default().Printf("[INFO] executed clean up of workdir: %s", workDirName)

}

func createWorkDir(baseName, preffix string) (string, error) {

	uuid := uuid.NewV4().String()
	var dirName = ""
	if preffix == "" {
		dirName = fmt.Sprintf("%s_%s", baseName, uuid)
	} else {
		dirName = fmt.Sprintf("%s_%s_%s", baseName, preffix, uuid)
	}

	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}

	log.Default().Println("[INFO] workdir was created")

	err = os.Chmod(dirName, 0777)
	if err != nil {
		return "", err
	}

	log.Default().Println("[INFO] chmod was applyed on workdir")

	return dirName, nil

}

func removeWorkDir(name string) error {
	err := os.RemoveAll(name)
	if err != nil {
		return err
	}
	return nil
}

//  validateRequestContent validates if the number of uploaded files = 1 and  its suffixes are allowed
func validatePayload(files []*multipart.FileHeader) error {

	if len(files) == 0 {
		return errors.New("no files uploaded in request")
	}

	if len(files) > 1 {
		return errors.New("more than 1 file was uploaded")
	}

	allowedSuffixes := []string{".zip"}
	for _, file := range files {
		for _, suffix := range allowedSuffixes {
			if strings.HasSuffix(strings.ToLower(file.Filename), suffix) {
				return nil
			}
		}
	}
	return errors.New("invalid suffix of uploaded file name. Allows only [.zip]")
}
