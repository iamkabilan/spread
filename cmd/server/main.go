package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/internal/api"
	"github.com/iamkabilan/spread/internal/node"
	"github.com/joho/godotenv"
)

func main() {
	router := api.Router()

	go node.MonitorNodes()

	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: ", err)
	}

	if err := database.Initialize(); err != nil {
		fmt.Println(err)
		return
	}
	defer database.GetDB().Close()

	var API_PORT = os.Getenv("API_PORT")
	log.Printf("Starting server on port %s", API_PORT)
	err = http.ListenAndServe(":"+API_PORT, router)
	if err != nil {
		fmt.Println("Error in starting the server, %w \n", err)
	}
}
