package gcs

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type StorageConnection struct {
	Client *storage.Client
}

// GetGCSClient gets singleton object for Google Storage
func NewStorageConnection() *StorageConnection {
	keyfilename := os.Getenv("KEY_FILE_NAME")
	if keyfilename == "" {
		keyfilename = "prodkeyfile.json"
	}
	keyfilepath := "/app/cert/" + keyfilename
	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile(keyfilepath))
	if err != nil {
		log.Printf("failed to create gcs client error:%s", err.Error())
		return nil
	}

	return &StorageConnection{
		Client: storageClient,
	}
}
