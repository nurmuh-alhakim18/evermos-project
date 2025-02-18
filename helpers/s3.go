package helpers

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

const (
	MaxFileSize      = 2 * 1024 * 1024 // 2 MB
	AllowedMimeTypes = "image/jpeg,image/png"
)

func LoadS3Session() {
	accessKeyID := GetEnv("AWS_ACCESS_KEY_ID", "")
	secretAccessKey := GetEnv("AWS_SECRET_ACCESS_KEY", "")
	region := GetEnv("AWS_REGION", "ap-southeast-3")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	S3Client = s3.NewFromConfig(cfg)
}

func UploadToS3(file *multipart.FileHeader) (string, error) {
	err := validateFile(file)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	shortID, err := GenerateShortID()
	if err != nil {
		return "", err
	}

	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%s%d%s", shortID, time.Now().Unix(), fileExt)
	filePath := fmt.Sprintf("uploads/%s", fileName)
	bucketName := GetEnv("BUCKET_NAME", "")
	contentType := file.Header.Get("Content-Type")

	_, err = S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &filePath,
		Body:        src,
		ACL:         "public-read",
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}

	region := GetEnv("AWS_REGION", "ap-southeast-3")
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, filePath)
	return fileURL, nil
}

func DeleteFromS3(url string) error {
	urlParts := strings.SplitN(url, "com/", 2)
	key := urlParts[1]
	bucketName := GetEnv("BUCKET_NAME", "")

	_, err := S3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	})

	if err != nil {
		return err
	}

	return nil
}

func validateFile(file *multipart.FileHeader) error {
	if file.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds %dMB", MaxFileSize)
	}

	contentType := file.Header.Get("Content-Type")
	types := strings.Split(AllowedMimeTypes, ",")
	if !slices.Contains(types, contentType) {
		return fmt.Errorf("invalid file type, allowed: jpeg, png")
	}

	return nil
}
