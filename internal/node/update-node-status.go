package node

import (
	"log"

	"github.com/iamkabilan/spread/database"
)

func UpdateNodeStatus(nodeID string, status string) error {
	db := database.GetDB()
	query := `UPDATE storage_nodes SET status = ?, last_heartbeat = CURRENT_TIMESTAMP WHERE node_id = ?`

	_, err := db.Exec(query, status, nodeID)
	if err != nil {
		log.Printf("Error updating node status for %s: %v", nodeID, err)
		return err
	}

	log.Printf("Node %s status updated to '%s'", nodeID, status)
	return nil
}
