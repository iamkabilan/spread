package controller

import (
	"fmt"
	"net/http"
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
		http.Error(response, "Error reading the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded file is %s", handler.Filename)
	fmt.Fprintf(response, "Uploaded file is %s", handler.Filename)
}
