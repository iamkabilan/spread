package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/internal/metadata"
	"github.com/iamkabilan/spread/internal/node"
	"github.com/joho/godotenv"
)

func main() {
	var wg sync.WaitGroup
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: ", err)
	}
	if err := database.Initialize(); err != nil {
		fmt.Println(err)
		return
	}
	defer database.GetDB().Close()

	port := flag.Int("port", 3000, "Port to start the node")
	flag.Parse()

	exists, nodeID, err := node.CheckIfNodeExists(*port)
	if err != nil {
		fmt.Printf("Error registering the node, %v", err)
		return
	}

	if !exists {
		log.Printf("Creating new node on port %d", *port)
		nodeID = uuid.New().String()
		metadata.SaveNewNode(nodeID, *port)
	} else {
		log.Printf("Node already exists on this port %d, setting it to active", *port)
		node.UpdateNodeStatus(nodeID, "active")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		node.SendHeartbeat(nodeID)
	}()

	wg.Wait()
}
