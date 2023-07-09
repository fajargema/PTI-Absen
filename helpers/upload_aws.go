package helpers

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(svc *s3.S3, file io.Reader, filename string) (string, error) {
	key := generateUniqueFileName(filename)

	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(file); err != nil {
		return "", err
	}

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("pti-absen"),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer.Bytes()),
	})
	if err != nil {
		return "", err
	}

	url := generateS3PublicURL("pti-absen", key)

	return url, nil
}

func generateUniqueFileName(filename string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	uniqueName := fmt.Sprintf("%s_%d%s", name, time.Now().UnixNano(), ext)
	return uniqueName
}

func generateS3PublicURL(bucket, key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
}
