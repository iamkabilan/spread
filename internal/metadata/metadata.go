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
	log.Println("New file inserted with the id", fileMetaData.FileId)

	return fileId, nil
}

func FetchFileMetaData(fileID int) (models.File, error) {
	var file models.File
	db := database.GetDB()
	query := "SELECT file_id, file_name, file_size, is_deleted FROM files WHERE file_id = ?"
	rows, err := db.Query(query, fileID)
	if err != nil {
		log.Printf("Error in fetching file metadata %v", err)
		return file, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&file.FileId, &file.FileName, &file.FileSize, &file.IsDeleted); err != nil {
			log.Printf("Error in reading rows from file table %v", err)
			return file, err
		}
	}

	return file, nil
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
