package api

import (
	"github.com/gorilla/mux"
	"github.com/iamkabilan/spread/internal/api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/upload", controller.UploadFile).Methods("POST")
	router.HandleFunc("/register-node", controller.RegisterNode).Methods("POST")
	router.HandleFunc("/download/{fileId}", controller.DownloadFile).Methods("GET")

	return router
}
