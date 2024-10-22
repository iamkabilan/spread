package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/internal/metadata"
	"github.com/iamkabilan/spread/models"
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
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return false
	}

	for chunkIndex, chunk := range chunks {
		chunkSize := len(chunk)
		chunkHash := calculateChunkHash(chunk)

		query := `INSERT INTO chunks (file_id, chunk_index, chunk_size, chunk_hash) VALUES (?, ?, ?, ?)`
		_, queryErr := db.Exec(query, fileID, chunkIndex, chunkSize, chunkHash)
		if queryErr != nil {
			log.Println("ERROR: ", queryErr.Error())

			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Println(rollbackErr)
			}
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

func StoreFile(fileBytes []byte, filename string, fileSize int64) (int64, error) {
	log.Println("File size is -----------------", fileSize)
	var file models.File
	file = models.File{
		UserId:   1,
		FileName: filename,
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
