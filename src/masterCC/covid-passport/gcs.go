package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

func ReadFileFromGCS(bucket, path string) ([]byte, error) {
	ctx, cnFnc := context.WithTimeout(context.Background(), time.Minute)
	defer cnFnc()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Creates a Bucket instance.
	buc := client.Bucket(bucket)
	obj := buc.Object(path)
	rdr, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error accessing GCS bucket: %w", err)
	}
	defer rdr.Close()
	data, err := io.ReadAll(rdr)
	if err != nil {
		return nil, fmt.Errorf("Error reading from GCS bucket: %w", err)
	}

	return data, nil
}
