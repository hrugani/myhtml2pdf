package services

import (
	"errors"
	"fmt"
	"log"
	"runtime"

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
	defer outFile.Close()
	errFile, err := os.Create("err.log")
	if err != nil {
		return "", err
	}
	defer errFile.Close()

	cmdName := "../wkhtmltopdf"
	if isWindows() {
		cmdName = "../wkhtmltopdf.exe"
		log.Default().Printf("Windows detected: comand name to be used: %s", cmdName)
	}

	cmd := exec.Command(cmdName, htmlFileName, "output.pdf")
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

func isWindows() bool {
	os := runtime.GOOS
	return os == "windows"
}
