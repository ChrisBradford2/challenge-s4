package models

import (
	"time"
)

type File struct {
	Base
	HackathonID  uint      `json:"hackathon_id"`                 // Foreign key for Hackathon
	UserID       string    `gorm:"index;not null" json:"userId"` // User ID
	Name         string    `json:"name"`                         // Filename
	Size         int64     `json:"size"`                         // File size in bytes
	ContentType  string    `json:"contentType"`                  // File content type
	LastModified time.Time `json:"lastModified"`                 // Last modified date
	URL          string    `json:"url"`                          // URL to download the file from Google Cloud Storage
}

type FileInfo struct {
	Name       string `json:"name"`
	UploadTime string `json:"upload_time"`
	URL        string `json:"url"`
}
