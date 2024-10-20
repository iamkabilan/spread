package models

type Chunk struct {
	ChunkId    int    `json:"chunk_id"`
	FileId     int    `json:"file_id"`
	ChunkIndex string `json:"chunk_index"`
	ChunkSize  string `json:"chunk_size"`
	ChunkHash  string `json:"chunk_hash"`
}
