package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/google/uuid"
	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/internal/metadata"
	"github.com/iamkabilan/spread/internal/node"
	pb "github.com/iamkabilan/spread/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
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

	go func() {
		node.SendHeartbeat(nodeID)
	}()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	chunkServer := &node.ChunkServer{}
	pb.RegisterChunkServiceServer(grpcServer, chunkServer)

	log.Printf("Node is listening on  port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
