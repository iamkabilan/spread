package storage

import (
	"log"

	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/models"
)

func RetrieveFile(file models.File) {
	db := database.GetDB()
	query := "SELECT chunk_id, chunk_index FROM chunks WHERE file_id = ?"
	rows, err := db.Query(query, file.FileId)
	if err != nil {
		log.Printf("Error in fetching the chunks. %v", err)
	}
	defer rows.Close()

	var chunks []models.Chunk
	for rows.Next() {
		var chunk models.Chunk
		if err := rows.Scan(&chunk.ChunkId, &chunk.ChunkIndex); err != nil {
			log.Printf("Error in reading the chunks, %v", err)
		}

		chunks = append(chunks, chunk)
	}
}
