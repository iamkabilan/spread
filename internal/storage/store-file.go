package storage

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/internal/metadata"
	"github.com/iamkabilan/spread/models"
	pb "github.com/iamkabilan/spread/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createChunks(fileBytes []byte) map[int][]byte {
	var chunks = make(map[int][]byte)
	var chunkIndex = 0
	chunkSize := 256 << 10 // Each chunk is 256 KB

	for i := 0; i < len(fileBytes); i += chunkSize {
		end := i + chunkSize

		if end > len(fileBytes) {
			end = len(fileBytes)
		}

		chunks[chunkIndex] = fileBytes[i:end]
		chunkIndex++
	}

	return chunks
}

func calculateChunkHash(chunk []byte) string {
	hash := sha256.New()
	hash.Write(chunk)
	return hex.EncodeToString(hash.Sum(nil))
}

func storeChunks(chunks map[int][]byte, fileID int64) bool {
	db := database.GetDB()

	activeNodes, err := getActiveNodes(db)
	if err != nil || len(activeNodes) == 0 {
		log.Printf("Error: %v", err)
		return false
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return false
	}

	nodeIndex := 0

	for chunkIndex, chunk := range chunks {
		chunkSize := len(chunk)
		chunkHash := calculateChunkHash(chunk)

		selectedNode := activeNodes[nodeIndex]
		nodeIndex = (nodeIndex + 1) % len(activeNodes)

		query := `INSERT INTO chunks (file_id, node_id, chunk_index, chunk_size, chunk_hash) VALUES (?, ?, ?, ?, ?)`
		result, queryErr := db.Exec(query, fileID, selectedNode.NodeID, chunkIndex, chunkSize, chunkHash)
		if queryErr != nil {
			log.Println("ERROR: ", queryErr.Error())

			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println(rollbackErr)
			}
			return false
		}

		chunkID, _ := result.LastInsertId()

		nodeResult := storeChunkOnNode(":"+strconv.Itoa(selectedNode.Port), fileID, chunkID, chunk)
		if nodeResult == false {
			log.Printf("Couldn't able to store the chunk.")
			tx.Rollback()
			return false
		}

		log.Printf("Chunk %d stored successfully.\n", chunkIndex)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		log.Println(commitErr)
		return false
	}
	return true
}

func storeChunkOnNode(nodeAddress string, fileID int64, chunkID int64, chunk []byte) bool {
	conn, err := grpc.NewClient(nodeAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	client := pb.NewChunkServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.StoreChunk(ctx, &pb.StoreChunkRequest{
		ChunkId: chunkID,
		FileId:  fileID,
		Chunk:   chunk,
	})
	if err != nil {
		log.Printf("Error: %v", err)
		return false
	}
	log.Printf(response.Message)
	return true
}

func StoreFile(fileBytes []byte, filename string, fileType string, fileSize int64) (int64, error) {
	var file models.File
	file = models.File{
		UserId:   1,
		FileName: filename,
		FileType: fileType,
		FileSize: fileSize,
	}

	fileID, err := metadata.SaveFileMetadata(file)
	if err != nil {
		fmt.Printf("Error in saving the file metadata, %v", err)
		return 0, err
	}

	chunks := createChunks(fileBytes)
	storeChunks(chunks, fileID)

	return fileID, nil
}
