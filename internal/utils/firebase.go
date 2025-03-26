package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

const (
	WASESSEN_FOLDER = "wasessen"
)

type FireBaseStorage struct {
	Bucket string
}

func UploadFile(ctx context.Context, file multipart.File) (string, error) {
	bucketName := os.Getenv("FIREBASE_BUCKET_NAME")
	if bucketName == "" {
		return "", errors.New("firebase bucket env variable empty")
	}

	fileName := uuid.New().String()

	fb := newFireBaseStorage(bucketName)

	creds, err := base64.StdEncoding.DecodeString(os.Getenv("FIREBASE_CREDENTIALS"))
	if err != nil {
		return "", fmt.Errorf("error decoding credentials: %v", err)
	}
	opt := option.WithCredentialsJSON(creds)

	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return "", fmt.Errorf("firebase client error: %v", err)
	}

	filePath := fmt.Sprintf("%s/%s", WASESSEN_FOLDER, fileName)
	wc := client.Bucket(fb.Bucket).Object(filePath).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("error copying file to storage: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("error closing storage writer: %v", err)
	}

	imageURL, err := generateFirebaseUrl(WASESSEN_FOLDER, fileName)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func newFireBaseStorage(bucket string) *FireBaseStorage {
	return &FireBaseStorage{
		Bucket: bucket,
	}
}

func generateFirebaseUrl(fileFolder, fileName string) (string, error) {
	baseURL := os.Getenv("FIREBASE_URL")
	if baseURL == "" {
		return "", errors.New("baseURL env variable empty")
	}

	if fileFolder == "" {
		return "", errors.New("file folder empty")
	}

	if fileName == "" {
		return "", errors.New("file name empty")
	}

	url := fmt.Sprintf("%s%s%%2F%s?alt=media", baseURL, fileFolder, fileName)

	return url, nil
}
