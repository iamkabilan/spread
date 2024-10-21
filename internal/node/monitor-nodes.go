package node

import (
	"log"
	"time"

	"github.com/iamkabilan/spread/database"
)

func MonitorNodes() {
	threshold := 30 * time.Second
	for {
		log.Println("Monitoring Nodes ....")

		db := database.GetDB()
		query := "SELECT node_id FROM storage_nodes WHERE TIMESTAMPDIFF(SECOND, last_heartbeat, NOW()) > ? AND status = 'active'"
		rows, queryErr := db.Query(query, int(threshold.Seconds()))
		if queryErr != nil {
			log.Printf("ERROR querying nodes: %v", queryErr.Error())
			continue
		}
		defer rows.Close()

		var nodeID string
		for rows.Next() {
			if err := rows.Scan(&nodeID); err != nil {
				log.Printf("Error scanning nodeID %v", err)
				continue
			}

			err := UpdateNodeStatus(nodeID, "down")
			log.Printf("Downing the node %s", nodeID)
			if err != nil {
				log.Printf("Error updating node %s to down: %v", nodeID, err)
			}
		}

		time.Sleep(15 * time.Second)
	}
}
