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
	log.Printf("%d", chunkID)

	return &pb.GetChunkResponse{
		Chunk:  nil,
		FileId: 1,
	}, nil
}

func storeChunkOnDisk(nodeAddress string, chunkID int64, fileID int64, chunk []byte) error {
	baseStoragePath := os.Getenv("BASE_STORAGE_PATH")
	log.Printf("%d %d %d", chunkID, fileID, len(chunk))
	filePath := filepath.Join(baseStoragePath, "file-storage", nodeAddress, fmt.Sprintf("file_%d-chunk_%d", fileID, chunkID))
	return os.WriteFile(filePath, chunk, 0644)
}
