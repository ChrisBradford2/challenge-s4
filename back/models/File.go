package models

import "time"

type File struct {
	Base
	UserID       string    `json:"userId"`       // User ID
	Name         string    `json:"name"`         // Filename
	Size         int64     `json:"size"`         // File size in bytes
	ContentType  string    `json:"contentType"`  // File content type
	LastModified time.Time `json:"lastModified"` // Last modified date
	URL          string    `json:"url"`          // URL to download the file from Google Cloud Storage
}
