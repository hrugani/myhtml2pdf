package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
)

// PdfMerger Merges all pdf files pointed by pdfFileNames sorting its names to define the sequence of appending process.
func MergePdfs(workDirName string, pdfFileNames []string) (string, error) {
	sortedPdfFileNames := sortPdfFileNames(pdfFileNames)
	outputFileName := "merged.pdf"
	// err := api.MergeAppendFile(sortedPdfFileNames, outputFileName, nil)
	err := mergePdfUsingPdfTk(sortedPdfFileNames, outputFileName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", workDirName, outputFileName), nil
}

func sortPdfFileNames(inputSlice []string) []string {
	sort.SliceStable(
		inputSlice,
		func(i, j int) bool { return inputSlice[i] < inputSlice[j] },
	)
	return inputSlice
}

func mergePdfUsingPdfTk(fNames []string, outputFileName string) error {

	outFile, err := os.Create("out.log")
	if err != nil {
		return err
	}
	defer outFile.Close()
	errFile, err := os.Create("err.log")
	if err != nil {
		return err
	}
	defer errFile.Close()

	// in linux the pdftk must be instlled via apt-get or snap
	cmdName := "pdftk"
	if isWindows() {
		// in windows the binary of pdftk must be present in parent directory
		cmdName = "../pdftk.exe"
		log.Default().Printf("[INFO] OS windows detected: comand name to be used: %s", cmdName)
	}

	var cmdLineParams []string
	cmdLineParams = append(cmdLineParams, fNames...)
	cmdLineParams = append(cmdLineParams, "cat")
	cmdLineParams = append(cmdLineParams, "output")
	cmdLineParams = append(cmdLineParams, outputFileName)
	cmd := exec.Command(cmdName, cmdLineParams...)
	cmd.Stdout = outFile
	cmd.Stderr = errFile
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		msgerr := fmt.Sprintf("[ERROR] executing pdftk command line. detail: %s", err)
		log.Default().Print(msgerr)
		return errors.New(msgerr)
	}

	log.Default().Print("[INFO] pdftk command-line executed")

	return nil

}
