package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func ConcatPdfs(workDirName, fName string) (string, error) {

	log.Default().Printf("[INFO] service ConcatPdfs was called. WorkDir= %s, File= %s", workDirName, fName)

	if err := changeCurrentDir(workDirName); err != nil {
		errMsg := fmt.Sprintf("[ERROR] changing current dir. detail: %s", err)
		log.Default().Print(errMsg)
		return "", err
	}
	defer changeCurrentDir("..")
	log.Default().Println("[INFO] current dir was changed to workdir")

	unzippedFilenames, err := Unzip(fName, ".")
	if err != nil {
		errMsg := fmt.Sprintf("[ERROR] unzipping file current dir. detail: %s", err)
		log.Default().Print(errMsg)
		return "", err
	}
	log.Default().Println("[INFO] Uploaded file was unzipped successfully")

	pdfFileNames, err := getPdfsFileNamesFromUnzippedFileNames(unzippedFilenames)
	if err != nil {
		errMsg := fmt.Sprintf("[ERROR] getting pdf file names. detail: %s", err)
		log.Default().Print(errMsg)
		return "", err
	}
	log.Default().Printf("[INFO] all pdf file names was identified: %s", pdfFileNames)

	pdfFileFullName, err := MergePdfs(workDirName, pdfFileNames)
	if err != nil {
		errMsg := fmt.Sprintf("[ERROR] merging pdf files. detail: %s", err)
		log.Default().Print(errMsg)
		return "", err
	}
	log.Default().Println("[INFO] pdfCpu command line util executed successful")
	log.Default().Println("[INFO] service ConcatPdfs finished successfully")
	return pdfFileFullName, nil
}

func getPdfsFileNamesFromUnzippedFileNames(unzippedFileNames []string) ([]string, error) {
	resp := []string{}
	for _, name := range unzippedFileNames {
		switch {
		case strings.HasSuffix(name, ".pdf"):
			resp = append(resp, name)
		}
	}
	if len(resp) <= 0 {
		return nil, errors.New("no pdf files in zipped uploaded file")
	}
	if len(resp) == 1 {
		return nil, errors.New("only 1 pdf file in zipped uploaded file")
	}
	return resp, nil
}
