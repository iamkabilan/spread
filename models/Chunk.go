package models

type Chunk struct {
	ChunkId    int64  `json:"chunk_id"`
	FileId     int64  `json:"file_id"`
	Port       int    `json:"port"`
	ChunkIndex int    `json:"chunk_index"`
	ChunkSize  string `json:"chunk_size"`
	ChunkHash  string `json:"chunk_hash"`
}
