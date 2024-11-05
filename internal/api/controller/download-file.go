package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/spread/internal/metadata"
)

func DownloadFile(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	fileID := variables["fileId"]

	fileIDInt, _ := strconv.Atoi(fileID)
	file, err := metadata.FetchFileMetaData(fileIDInt)
	if err != nil {
		http.Error(response, "Couldn't able to download the file", http.StatusInternalServerError)
		return
	}

	log.Print(file)
	json.NewEncoder(response).Encode(file)
}
