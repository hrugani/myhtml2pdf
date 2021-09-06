package services

import (
	"fmt"
	"sort"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// PdfMerger Merges all pdf files pointed by pdfFileNames sorting its names to define the sequence of appending process.
func MergePdfs(workDirName string, pdfFileNames []string) (string, error) {
	sortedPdfFileNames := sortPdfFileNames(pdfFileNames)
	outputFileName := "output.pdf"
	err := api.MergeAppendFile(sortedPdfFileNames, outputFileName, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", workDirName, outputFileName), nil
}

func sortPdfFileNames(inputSlice []string) []string {
	sort.SliceStable(
		inputSlice,
		func(i, j int) bool { return inputSlice[i] > inputSlice[j] },
	)
	return inputSlice
}
