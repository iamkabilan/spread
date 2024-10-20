package storage

import (
	"fmt"

	"github.com/iamkabilan/spread/internal/metadata"
	"github.com/iamkabilan/spread/models"
)

func createChunks(fileBytes []byte) [][]byte {
	var chunks [][]byte
	chunkSize := 256 << 18 // Each chunk is 256 KB

	for i := 0; i < len(fileBytes); i += chunkSize {
		end := i + chunkSize

		if end > len(fileBytes) {
			end = len(fileBytes)
		}

		chunks = append(chunks, fileBytes[i:end])
	}

	return chunks
}

func StoreFile(fileBytes []byte, filename string, fileSize int64) (int64, error) {
	var file models.File
	file = models.File{
		UserId:   1,
		FileName: filename,
		FileSize: fileSize,
	}

	fileId, err := metadata.SaveFileMetadata(file)
	if err != nil {
		fmt.Printf("Error in saving the file metadata, %v", err)
		return 0, err
	}

	return fileId, nil

	// chunks := createChunks(fileBytes)

}
