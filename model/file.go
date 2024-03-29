package model

import "time"

type File struct {
	ID        string    `gorm:"primaryKey" json:"id_file"`
	FileName  string    `gorm:"not null;type:varchar(255)" json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FileResponse struct {
	ID        string    `json:"id_file"`
	FileName  string    `json:"file_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FileRequest struct {
	ID       string `json:"id_file"`
	FileName string `json:"file_name"`
}
