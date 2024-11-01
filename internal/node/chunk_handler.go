package node

import (
	"context"
	"log"

	pb "github.com/iamkabilan/spread/proto"
)

type ChunkServer struct {
	pb.UnimplementedChunkServiceServer
}

func (s *ChunkServer) StoreChunk(ctx context.Context, req *pb.StoreChunkRequest) (*pb.StoreChunkResponse, error) {
	chunkID := req.GetChunkId()
	fileID := req.GetFileId()
	chunk := req.GetChunk()

	log.Printf("Storing chunk id %d for file id %d", chunkID, fileID)
	err := storeChunkOnDisk(chunkID, fileID, chunk)
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

func storeChunkOnDisk(chunkID int64, fileID int64, chunk []byte) error {
	log.Printf("%d %d %d", chunkID, fileID, len(chunk))
	return nil
}
