package node

import (
	"log"
	"os"
	"path/filepath"

	"github.com/iamkabilan/spread/database"
)

var baseStoragePath = os.Getenv("BASE_STORAGE_PATH")

func InitNodeFolder(nodeAddress string) error {
	dir := filepath.Join(baseStoragePath, "file-storage", nodeAddress)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
		log.Printf("Node folder intialized %s", dir)
	}
	return nil
}

func CheckIfNodeExists(port int) (bool, string, error) {
	db := database.GetDB()
	query := "SELECT node_id FROM storage_nodes WHERE port = ?"
	rows, err := db.Query(query, port)
	if err != nil {
		log.Printf("Error Fetching the nodes, %v", err)
		return false, "", err
	}
	defer rows.Close()

	if rows.Next() {
		var nodeID string
		if err := rows.Scan(&nodeID); err != nil {
			log.Printf("Error scanning node ID: %v", err)
			return false, "", nil
		}
		return true, nodeID, nil
	}

	return false, "", nil
}
