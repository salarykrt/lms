package integrations

import (
	configuration "salarykart/config"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadDocument(filePath, objectName string, optionalParams ...string) error {
	s3Config := configuration.GlobalConfig.S3Bucket

	if _, err := os.Stat(filePath); err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Error opening file")
		return err
	}
	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Config.Region),
		Credentials: credentials.NewStaticCredentials(s3Config.AccessKey, s3Config.SecretKey, ""),
	})
	if err != nil {
		slog.Error("Error creating AWS session")
		return err
	}

	var UploadPath string
	// Check if optionalParams has been provided
	if len(optionalParams) > 0 {
		UploadPath = optionalParams[0]
	}

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:       aws.String(s3Config.Bucket + UploadPath),
		Key:          aws.String(objectName),
		Body:         file,
		StorageClass: aws.String(s3Config.StorageClass),
	})

	if err != nil {
		slog.Error("Error uploading document to S3")
		return err
	}

	return nil
}

func DownloadDocument(objectName, downloadPath string) error {
	s3Config := configuration.GlobalConfig.S3Bucket
	// If downloadPath is not provided, download the document to the temp_doc_path
	if downloadPath == "" {
		downloadPath = configuration.GlobalConfig.Local.TempDocPath + strings.Split(objectName, "/")[len(strings.Split(objectName, "/"))-1]
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Config.Region),
		Credentials: credentials.NewStaticCredentials(s3Config.AccessKey, s3Config.SecretKey, ""),
	})
	if err != nil {
		slog.Error("Error creating AWS session")
		return err
	}

	downloader := s3manager.NewDownloader(sess)
	file, err := os.Create(downloadPath)
	if err != nil {
		slog.Error("Error creating file")
		return err
	}
	defer file.Close()

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s3Config.Bucket),
			Key:    aws.String(objectName),
		})
	if err != nil {
		slog.Error("Error downloading document from S3")
		return err
	}

	return nil
}
