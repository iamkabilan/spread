package api

import (
	"github.com/gorilla/mux"
	"github.com/iamkabilan/spread/internal/api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/upload", controller.UploadFile).Methods("POST")

	return router
}
