package services

import (
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
