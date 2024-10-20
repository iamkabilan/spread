package node

import (
	"database/sql"
	"log"
)

func UpdateNodeStatus(db *sql.DB, nodeID string, status string) error {
	query := `UPDATE storage_node SET status = ?, last_heartbeat = CURRENT_TIMESTAMP WHERE node_id = ?`

	_, err := db.Exec(query, status, nodeID)
	if err != nil {
		log.Printf("Error updating node status for %s: %v", nodeID, err)
		return err
	}

	log.Printf("Node %s status updated to '%s'", nodeID, status)
	return nil
}
