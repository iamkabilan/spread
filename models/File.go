package models

type File struct {
	FileId    int    `json:"file_id"`
	UserId    int    `json:"user_id"`
	FileName  string `json:"filename"`
	FileSize  int64  `json:"file_size"`
	Version   int    `json:"version"`
	IsDeleted bool   `json:"is_deleted"`
}
