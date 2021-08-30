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
		c.String(
			http.StatusBadRequest,
			fmt.Sprintf("err: %s", err.Error()+"\n"),
		)
		return
	}
	files := form.File["files"]

	err = validatePayload(files)
	if err != nil {
		log.Default().Println(err)
		c.String(
			http.StatusBadRequest,
			fmt.Sprintf("err: %s", err.Error()+"\n"),
		)
		return
	}

	// creates workdir
	workDirName, err := createWorkDir()
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("err creating local work dir: %s", err.Error()+"\n"),
		)
		return
	}

	var uploadedFileName string
	for _, file := range files {

		// fmt.Printf("file received:  %#v \n\n", file)
		// c.String(
		// 	http.StatusOK,
		// 	fmt.Sprintf("file received:  %#v \n\n", file),
		// )
		pathSep := fmt.Sprintf("%c", os.PathSeparator)
		uploadedFileName = filepath.Base(file.Filename)
		fileNameInWorkDir := workDirName + pathSep + file.Filename
		if err := c.SaveUploadedFile(file, fileNameInWorkDir); err != nil {
			c.String(
				http.StatusBadRequest,
				fmt.Sprintf("upload file err: %s", err.Error()+"\n"),
			)
			return
		}
	}

	log.Default().Printf("file %v was uploaded and saved properly in the server", uploadedFileName)

	// process the html to pdf convertion
	var pdfFilePath string
	if strings.HasSuffix(strings.ToLower(uploadedFileName), ".zip") {
		pdfFilePath, err = services.Zip2Pdf(workDirName, uploadedFileName)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("server error: %s \n", err.Error()))
			return
		}
	} else {
		pdfFilePath, err = services.HtmlText2Pdf(workDirName, uploadedFileName)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("server err: %s", err.Error()+"\n"))
			return
		}
	}

	// reponse: pdf file
	c.File(string(pdfFilePath))
	log.Default().Printf("html to pdf successful. pdf file generated: %s", pdfFilePath)

	// Removes workdir
	err = removeWorkDir(workDirName)
	if err != nil {
		log.Default().Printf("err on deliting used workdir %s", workDirName)
	}
	log.Default().Printf("executed clean up of workdir: %s", workDirName)

}

func createWorkDir() (string, error) {

	uuid := uuid.NewV4().String()

	err := os.Mkdir(uuid, 0777)
	if err != nil {
		return "", err
	}

	err = os.Chmod(uuid, 0777)
	if err != nil {
		return "", err
	}
	
	return uuid, nil

	// // cmd := fmt.Sprintf("mkdir %s", uuid)
	// // out, err := exec.Command(cmd).Output()
	// // if err != nil {
	// // 	return "", err
	// // }
	// cmdOut := string(out)
	// _ = cmdOut
	// return uuid, nil
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

	allowedSuffixes := []string{".zip", ".html"}
	for _, file := range files {
		for _, suffix := range allowedSuffixes {
			if strings.HasSuffix(strings.ToLower(file.Filename), suffix) {
				return nil
			}
		}
	}
	return errors.New("invalid suffix of uploaded file name. Allows [.zip, .html]")
}
