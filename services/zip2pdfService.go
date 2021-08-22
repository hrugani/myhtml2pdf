package services

import "log"


func Zip2Pdf(workdirName, filename string) ([]byte, error) {
	log.Default().Printf("service Zip2Pdf was called. WorkDir= %s, File= %s", workdirName, filename )
	return nil, nil
}
