package services

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
)

func imagesEmbedder(htmlFileName string, imageFileNames []string) (string, error) {

	content, err := os.ReadFile(htmlFileName)
	if err != nil {
		return "", err
	}
	html := string(content)

	imgTags, err := getAllImgTags(&html)
	if err != nil {
		return "", err
	}

	replaceMap, err := getReplaceMap(imgTags, imageFileNames)
	if err != nil {
		return "", nil
	}

	for imgTag, imgBase64 := range replaceMap {
		html = embedImageInHtml(imgTag, imgBase64, html)
	}

	return "", nil
}

func getAllImgTags(html *string) ([]*string, error) {
	//todo
	return nil, nil
}

func embedImageInHtml(imgTag *string, imgBase64 *string, html string) string {
	//todo
	return ""
}

func getReplaceMap(imgTags []*string, imagesFileNames []string) (map[*string]*string, error) {

	imgTagImgFileNameMap, err := getImgTagImgFileNameMap(imgTags, imagesFileNames)
	if err != nil {
		return nil, err
	}

	resp := make(map[*string]*string)
	for imgTag, fileName := range imgTagImgFileNameMap {
		resp[imgTag] = calcImageBase64String(fileName)
	}

	return resp, nil
}

func getImgTagImgFileNameMap(imgTags []*string, imageFileNames []string) (map[*string]string, error) {
	return nil, nil
}

func calcImageBase64String(fName string) *string {

	bytes, err := os.ReadFile(fName)
	if err != nil {
		log.Default().Printf("error reading file: %s", fName)
	}


	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	return &base64Encoding
}

// toBase64 return the base64 string representation of the []byte
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
