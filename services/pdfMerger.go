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

	// builds the list od pdf files as it must be passed to the pdfth command-line application
	var pdfFileNames = " "
	for _, name := range fNames {
		pdfFileNames += name + " "
	}
	lastParam := fmt.Sprintf("cat output %s", outputFileName)

	log.Default().Printf("[INFO] executing command line: %s %s %s", cmdName, pdfFileNames, lastParam)

	cmd := exec.Command(cmdName, pdfFileNames, lastParam)
	cmd.Stdout = outFile
	cmd.Stderr = errFile
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		msgerr := fmt.Sprintf("[ERROR] executing command line. detail: %s", err)
		log.Default().Print(msgerr)
		return errors.New(msgerr)
	}

	log.Default().Print("[INFO] command line executed")

	return nil

}
