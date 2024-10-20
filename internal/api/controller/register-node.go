package controller

import (
	"flag"
	"net/http"

	"github.com/google/uuid"
	"github.com/iamkabilan/spread/internal/metadata"
)

func RegisterNode(response http.ResponseWriter, request *http.Request) {

	nodeID := uuid.New().String()
	port := flag.Int("port", 3000, "Port to start the node")

	metadata.SaveNewNode(nodeID, *port)
}
