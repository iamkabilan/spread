package metadata

import (
	"fmt"
	"log"

	"github.com/iamkabilan/spread/database"
	"github.com/iamkabilan/spread/models"
)

func SaveFileMetadata(fileMetaData models.File) (int64, error) {
	db := database.GetDB()
	query := "INSERT INTO files (user_id, filename, file_size) VALUES (?, ?, ?)"
	result, queryErr := db.Exec(query, fileMetaData.UserId, fileMetaData.FileName, fileMetaData.FileSize)
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr.Error())
		return 0, queryErr
	}

	fileId, _ := result.LastInsertId()
	fileMetaData.FileId = int(fileId)
	log.Println("File inserted ", fileMetaData.FileId)

	return fileId, nil
}

func SaveNewNode(nodeID string, port int) bool {
	db := database.GetDB()
	query := "INSERT INTO storage_nodes (node_id, port, location, status) VALUES (?, ?, ?, ?)"
	_, queryErr := db.Exec(query, nodeID, port, "earth", "active")
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr.Error())
		return false
	}
	log.Println("Stoarge node registered and it is active.")
	return true
}
