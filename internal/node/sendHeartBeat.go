package node

import (
	"log"
	"time"

	"github.com/iamkabilan/spread/database"
)

func SendHeartbeat(nodeID string) {
	db := database.GetDB()
	query := `UPDATE storage_nodes SET last_heartbeat = NOW(), status = 'active' where node_id = ?`

	for {
		_, queryErr := db.Exec(query, nodeID)
		if queryErr != nil {
			log.Println("ERROR: ", queryErr)
			continue
		}

		log.Printf("Sending heartbeat from the node %s", nodeID)
		time.Sleep(30 * time.Second)
	}
}
