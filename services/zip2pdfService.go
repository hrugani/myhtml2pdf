package services

import (
	"errors"
	"log"
	"os"
	"strings"
)

func Zip2Pdf(workdirName, filename string) (string, error) {

	log.Default().Printf("service Zip2Pdf was called. WorkDir= %s, File= %s", workdirName, filename)

	if err := changeCurrentDir(workdirName); err != nil {
		return "", err
	}
	defer changeCurrentDir("..")

	if err := unzip(filename); err != nil {
		return "", err
	}

	htmlFileName, err := embedImgsInHtml()
	if err != nil {
		return "", err
	}

	pdfFileFullName, err := convert(htmlFileName)
	if err != nil {
		return "", err
	}

	return pdfFileFullName, nil
}

func unzip(fName string) error {
	return nil
}

func changeCurrentDir(dirName string) error {
	if err := os.Chdir(dirName); err != nil {
		return err
	}
	return nil
}

func embedImgsInHtml() (string, error) {

	file, err := os.Open(".")
	if err != nil {
		return "", err
	}
	defer file.Close()

	htmlFileName, err := getHtmlFileName()
	if err != nil {
		return "", err
	}

	imageFileNames, err := getImageFileNames()
	if err != nil {
		return "", err
	}

	newHtmlFileName, err := imagesEmbeder(htmlFileName, imageFileNames)
	if err != nil {
		return "", err
	}

	return newHtmlFileName, nil
}

func getHtmlFileName() (string, error) {

	file, err := os.Open(".")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileList, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range fileList {
		if strings.HasSuffix(name, ".html") {
			return name, nil
		}
	}

	return "", errors.New("no html file present in uploaded data")

}

func getImageFileNames() ([]string, error) {
	file, err := os.Open(".")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileList, _ := file.Readdirnames(0) // 0 to read all files and folders
	resp := []string{}
	for _, name := range fileList {
		if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") {
			resp = append(resp, name)
		}
	}

	return resp, nil
}

func imagesEmbeder(htmlFileName string, imageFileNames []string) (string, error) {
	return "", nil
}

func convert(htmlFileName string) (string, error) {
	return "", nil
}
