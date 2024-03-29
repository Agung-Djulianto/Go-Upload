package helper

import (
	"log"
	"mime/multipart"

	"github.com/google/uuid"
)

func GenerateID() string {
	id := uuid.New()
	return id.String()
}

func IsValidImageFile(header *multipart.FileHeader) bool {
	allowedImageTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	contentType := header.Header.Get("Content-Type")

	isValid := allowedImageTypes[contentType]

	if !isValid {
		log.Printf("Invalid File Type. Allowed types: %v", allowedImageTypes)
	}

	return isValid
}
