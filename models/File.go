package models

type File struct {
	FileId    int    `json:"file_id"`
	UserId    int    `json:"user_id"`
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileSize  int64  `json:"file_size"`
	Version   int    `json:"version"`
	IsDeleted bool   `json:"is_deleted"`
}
