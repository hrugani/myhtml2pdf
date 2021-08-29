package services

import (
	"fmt"
	"os/exec"
)

func wkhtmltopdfConvert(workDirName, htmlFileName string) (string, error) {

	// out1, err1 := exec.Command("pwd").Output()
	// if err1 != nil {
	// 	return "", err1
	// }
	// cmd1 := string(out1)
	// _ = cmd1

	cmd := fmt.Sprintf("./wkhtmltopdf %s/%s %s/output.pdf", workDirName, htmlFileName, workDirName)
	//cmd := fmt.Sprintf("../wkhtmltopdf %s output.pdf", htmlFileName)
	out, err := exec.Command(cmd).Output()
	cmdOut := string(out)
	_ = cmdOut
	if err != nil {
		return "", err
	}

	resp := fmt.Sprintf("%s/output.pdf", workDirName)
	return resp, nil
}
