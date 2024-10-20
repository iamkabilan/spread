package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/iamkabilan/spread/internal/storage"
)

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

	fmt.Printf("Uploaded file is %s", handler.Filename)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(response, "Couldn't able to read the file.", http.StatusInternalServerError)
		return
	}
	log.Println("File bytes size is", len(fileBytes))

	fileId, err := storage.StoreFile(fileBytes, handler.Filename, handler.Size)
	if err != nil || fileId == 0 {
		http.Error(response, "Couldn't able to store the file.", http.StatusInternalServerError)
	}

	fmt.Fprintf(response, "Uploaded file is %s \n", handler.Filename)
}
