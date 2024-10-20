package models

type Node struct {
	NodeID        string `json:"node_id"`
	Port          int    `json:"port"`
	Location      string `json:"location"`
	Status        string `json:"status"`
	LastHeartbeat string `json:"last_heartbeat"`
}
