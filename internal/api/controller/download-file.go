package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/spread/internal/metadata"
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

	json.NewEncoder(response).Encode(fileMetadata)
}
