package services

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"log"
	"time"
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

func GenerateSignedURL(bucketName, objectName string, storageClient *storage.Client) (string, error) {
	urlValidFor := time.Duration(15) * time.Minute

	opts := &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(urlValidFor),
	}

	url, err := storageClient.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		return "", err
	}
	return url, nil
}
