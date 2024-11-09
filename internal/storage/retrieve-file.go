package storage

import (
	"bytes"
	"context"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/models"
	pb "github.com/iamkabilan/spread/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getChunkFromNode(nodeAddress string, fileID int64, chunkID int64) ([]byte, error) {
	conn, err := grpc.NewClient(nodeAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewChunkServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetChunk(ctx, &pb.GetChunkRequest{
		ChunkId: chunkID,
		FileId:  fileID,
	})

	if err != nil {
		return nil, err
	}

	return response.Chunk, nil
}

func writeChunksToFile(chunks map[int][]byte) []byte {
	fileBuffer := bytes.NewBuffer(nil)

	indices := make([]int, 0, len(chunks))
	for index := range chunks {
		indices = append(indices, index)
	}
	sort.Ints(indices)

	for _, chunkIndex := range indices {
		log.Printf("Writing chunk index %d", chunkIndex)
		fileBuffer.Write(chunks[chunkIndex])
	}

	return fileBuffer.Bytes()
}

func RetrieveFile(file models.File) ([]byte, error) {
	db := database.GetDB()
	query := "SELECT chunk_id, file_id, chunk_index, port FROM chunks c JOIN storage_nodes sn ON c.node_id = sn.node_id WHERE file_id = ?"
	rows, err := db.Query(query, file.FileId)
	if err != nil {
		log.Printf("Error in fetching the chunks. %v", err)
		return nil, err
	}
	defer rows.Close()

	var chunks = make(map[int][]byte)
	for rows.Next() {
		var chunk models.Chunk
		if err := rows.Scan(&chunk.ChunkId, &chunk.FileId, &chunk.ChunkIndex, &chunk.Port); err != nil {
			log.Printf("Error in reading the chunks, %v", err)
			return nil, err
		}

		chunkData, err := getChunkFromNode(":"+strconv.Itoa(chunk.Port), chunk.FileId, chunk.ChunkId)
		if err != nil {
			log.Printf("Error in establishing gRPC connection %v", err)
			return nil, err
		}

		log.Print(len(chunkData))
		chunks[chunk.ChunkIndex] = chunkData
	}

	fileBuffer := writeChunksToFile(chunks)

	return fileBuffer, nil
}
