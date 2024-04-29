package services

import (
	"mime/multipart"
	"net/http"
)

func IsValidImageType(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return false
	}

	contentType := http.DetectContentType(buf)
	switch contentType {
	case "image/jpeg", "image/png":
		return true
	default:
		return false
	}
}
