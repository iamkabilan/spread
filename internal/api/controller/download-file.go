package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/spread/internal/metadata"
	"github.com/iamkabilan/spread/internal/storage"
)

func DownloadFile(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	fileID := variables["fileId"]

	fileIDInt, _ := strconv.Atoi(fileID)
	fileMetadata, err := metadata.FetchFileMetaData(fileIDInt)
	if err != nil {
		http.Error(response, "Couldn't able to download the file", http.StatusInternalServerError)
		return
	}

	if fileMetadata.FileId == 0 || fileMetadata.IsDeleted {
		http.Error(response, "File not found", http.StatusNotFound)
		return
	}

	file, err := storage.RetrieveFile(fileMetadata)
	if err != nil {
		log.Printf("Error in retrieving the file, %v", err)
		http.Error(response, "Error in retrieving the file", http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", fileMetadata.FileType)
	response.Header().Set("Content-Disposition", "attachment; filename="+fileMetadata.FileName)
	response.Write(file)
}
