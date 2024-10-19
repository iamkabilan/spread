package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iamkabilan/spread/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	router := api.Router()

	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: ", err)
	}
	var API_PORT = os.Getenv("API_PORT")

	log.Printf("Starting server on port %s", API_PORT)
	http.ListenAndServe(":"+API_PORT, router)
}
