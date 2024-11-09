package controller

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/iamkabilan/spread/internal/storage"
)

func findContentType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error in reading the file %v", err)
		return "", err
	}
	mimeType := http.DetectContentType(buffer)
	log.Printf("File type %s", mimeType)

	return mimeType, nil
}

func UploadFile(response http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(10 << 20) // Allowing file size upto 10 MB

	if err != nil {
		http.Error(response, "File is too large. Maximum allowed limit is 10 MB.", http.StatusRequestEntityTooLarge)
		return
	}

	file, handler, err := request.FormFile("file")
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(response, "Error uploading the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded file is %s \n", handler.Filename)

	fileType, err := findContentType(file)
	if err != nil {
		http.Error(response, "Unable to find the file type", http.StatusBadRequest)
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(response, "Couldn't able to read the file.", http.StatusInternalServerError)
		return
	}
	log.Println("File bytes size is", len(fileBytes))

	var wg sync.WaitGroup
	wg.Add(1)

	go func(fileBytes []byte, filename string, fileType string, size int64) {
		defer wg.Done()

		fileId, err := storage.StoreFile(fileBytes, filename, fileType, size)
		if err != nil || fileId == 0 {
			http.Error(response, "Couldn't able to store the file.", http.StatusInternalServerError)
			return
		}
	}(fileBytes, handler.Filename, fileType, handler.Size)
	wg.Wait()

	fmt.Fprintf(response, "Uploaded file is %s \n", handler.Filename)
}
