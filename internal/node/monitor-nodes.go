package node

import (
	"fmt"
	"log"
	"time"

	"github.com/iamkabilan/spread/database"
)

func MonitorNodes() {
	for {
		threshold := 30 * time.Second

		db := database.GetDB()
		query := "SELECT node_id FROM storage_node WHERE TIMESTAMPDIFF(SECOND, last_heartbeat, NOW()) > ? AND status = 'active'"
		rows, queryErr := db.Query(query, int(threshold))
		if queryErr != nil {
			fmt.Println("ERROR querying nodes: ", queryErr.Error())
			continue
		}
		defer rows.Close()

		var nodeID string
		for rows.Next() {
			if err := rows.Scan(&nodeID); err != nil {
				log.Printf("Error scanning nodeID %v", err)
				continue
			}

			err := UpdateNodeStatus(db, nodeID, "down")
			if err != nil {
				log.Printf("Error updating node %s to down: %v", nodeID, err)
			}
		}

		time.Sleep(15 * time.Second)
	}
}
