package node

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	pb "github.com/iamkabilan/spread/proto"
)

type ChunkServer struct {
	pb.UnimplementedChunkServiceServer
	Port string
}

func (s *ChunkServer) StoreChunk(ctx context.Context, req *pb.StoreChunkRequest) (*pb.StoreChunkResponse, error) {
	chunkID := req.GetChunkId()
	fileID := req.GetFileId()
	chunk := req.GetChunk()

	log.Printf("Storing chunk id %d for file id %d", chunkID, fileID)
	err := storeChunkOnDisk(s.Port, chunkID, fileID, chunk)
	if err != nil {
		return &pb.StoreChunkResponse{
			Success: false,
			Message: "Failed to store the chunk",
		}, err
	}

	return &pb.StoreChunkResponse{
		Success: true,
		Message: "Successfully stored the chunk",
	}, nil
}

func (s *ChunkServer) GetChunk(ctx context.Context, req *pb.GetChunkRequest) (*pb.GetChunkResponse, error) {
	chunkID := req.GetChunkId()
	fileID := req.GetFileId()
	log.Printf("%d", chunkID)

	chunk, err := retrieveChunkFromDisk(s.Port, chunkID, fileID)
	if err != nil {
		log.Printf("Error in reading file from the storage %v", err)
		return &pb.GetChunkResponse{
			Chunk: nil,
		}, err
	}

	return &pb.GetChunkResponse{
		Chunk: chunk,
	}, nil
}

func storeChunkOnDisk(nodeAddress string, chunkID int64, fileID int64, chunk []byte) error {
	baseStoragePath := os.Getenv("BASE_STORAGE_PATH")
	filePath := filepath.Join(baseStoragePath, "file-storage", nodeAddress, fmt.Sprintf("file_%d-chunk_%d", fileID, chunkID))
	return os.WriteFile(filePath, chunk, 0644)
}

func retrieveChunkFromDisk(nodeAddress string, chunkID int64, fileID int64) ([]byte, error) {
	baseStoragePath := os.Getenv("BASE_STORAGE_PATH")
	filePath := filepath.Join(baseStoragePath, "file-storage", nodeAddress, fmt.Sprintf("file_%d-chunk_%d", fileID, chunkID))

	chunk, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}
