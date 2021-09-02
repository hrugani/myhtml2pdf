package services

import (
	"archive/zip"
	"errors"

	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Zip2Pdf executes the main peocess that converts a html file to a pdf file
func Zip2Pdf(workdirName, filename string) (string, error) {

	log.Default().Printf("[INFO] service Zip2Pdf was called. WorkDir= %s, File= %s", workdirName, filename)

	if err := changeCurrentDir(workdirName); err != nil {
		return "", err
	}
	defer changeCurrentDir("..")
	log.Default().Println("[INFO] current dir was changed to workdir")

	unzippedFilenames, err := Unzip(filename, ".")
	if err != nil {
		return "", err
	}
	log.Default().Println("[INFO] Uploaded file was unzipped successfully")

	htmlFileName, err := getHtmlFileNameFromUnzipedFileNames(unzippedFilenames)
	if err != nil {
		return "", err
	}
	log.Default().Printf("[INFO] html file name was identified: %s", htmlFileName)

	imagesFileNames := getImageFileNamesFromUnzipdFileNames(unzippedFilenames)
	log.Default().Printf("[INFO] all image file names was identified, %v ", imagesFileNames)

	// only if html contains images embed them.
	if len(imagesFileNames) > 0 {
		htmlFileName, err = imagesEmbedder(htmlFileName, imagesFileNames)
		if err != nil {
			return "", err
		}
	}

	pdfFileFullName, err := wkhtmltopdfConvert(workdirName, htmlFileName)
	if err != nil {
		return "", err
	}
	log.Default().Println("[INFO] wkhtmltopdf command line util executed successful")
	log.Default().Println("[INFO] service Zip2Pdf finished successfully")
	return pdfFileFullName, nil
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		// if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
		// 	return filenames, fmt.Errorf("%s: illegal file path", fpath)
		// }

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func changeCurrentDir(dirName string) error {
	if err := os.Chdir(dirName); err != nil {
		return err
	}
	return nil
}

func getHtmlFileNameFromUnzipedFileNames(fNames []string) (string, error) {
	for _, name := range fNames {
		if strings.HasSuffix(name, ".html") {
			return name, nil
		}
	}
	return "", errors.New("no html file don't present in uploaded data")
}

func getImageFileNamesFromUnzipdFileNames(fNames []string) []string {
	resp := []string{}
	for _, name := range fNames {
		switch {
		case strings.HasSuffix(name, ".png"):
			resp = append(resp, name)
		case strings.HasSuffix(name, ".jpg"):
			resp = append(resp, name)
		case strings.HasSuffix(name, ".jpeg"):
			resp = append(resp, name)
		case strings.HasSuffix(name, ".gif"):
			resp = append(resp, name)
		}
	}
	return resp
}
