package services

import "log"

func HtmlText2Pdf(workDirName, fName string) (string, error) {
	log.Default().Printf("service Zip2Pdf was called. WorkDir= %s, File= %s", workDirName, fName)
	return "", nil
}
