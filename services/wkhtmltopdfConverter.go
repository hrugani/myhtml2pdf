package services

import (
	"errors"
	"fmt"

	// "log"
	"os"
	"os/exec"
	// "strings"
)

func wkhtmltopdfConvert(workDirName, htmlFileName string) (string, error) {

	outFile, err := os.Create("out.log")
	if err != nil {
		return "", err
	}
	errFile, err := os.Create("err.log")
	if err != nil {
		return "", err
	}
	cmd := exec.Command("../wkhtmltopdf", htmlFileName, "output.pdf")
	cmd.Stdout = outFile
	cmd.Stderr = errFile
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		msgerr := "[ERROR] probably some image isn't present into uloaded files"
		return "", errors.New(msgerr)
	}

	resp := fmt.Sprintf("%s/output.pdf", workDirName)
	return resp, nil
}
