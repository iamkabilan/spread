package storage

import (
	"database/sql"
	"log"

	"github.com/iamkabilan/spread/models"
)

func getActiveNodes(db *sql.DB) ([]models.Node, error) {
	query := "SELECT node_id, port FROM storage_nodes WHERE status = 'active'"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer rows.Close()

	var nodes []models.Node
	for rows.Next() {
		var node models.Node
		if err := rows.Scan(&node.NodeID, &node.Port); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
