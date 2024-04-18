package services

import (
	"challenges4/models"
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"log"
)

type StorageService struct {
	Client *storage.Client
}

func NewStorageService(ctx context.Context, credentialsFile string) *StorageService {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return &StorageService{
		Client: client,
	}
}

func (s *StorageService) GetFilesByUserID(ctx context.Context, userID string) ([]models.File, error) {
	var files []models.File
	bucketName := "my-bucket"
	query := &storage.Query{
		Prefix: userID + "/",
	}
	it := s.Client.Bucket(bucketName).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err != nil {
			break
		}
		files = append(files, models.File{
			UserID:       userID,
			Name:         attrs.Name,
			Size:         attrs.Size,
			ContentType:  attrs.ContentType,
			LastModified: attrs.Updated,
			URL:          "https://storage.googleapis.com/" + bucketName + "/" + attrs.Name,
		})
	}

	return files, nil
}
