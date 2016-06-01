package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Config struct {
	AwsAccessKeyID     string `toml:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `toml:"AWS_SECRET_ACCESS_KEY"`
	AwsS3BucketName    string `toml:"AWS_S3_BUCKET_NAME"`
	AwsS3Region        string `toml:"AWS_S3_REGION"`
}

type S3 struct {
	S3Config
}

func NewS3(configPath string) *S3 {
	conf := S3Config{}
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &S3{
		S3Config: conf,
	}
}

func (s *S3) Upload(f *os.File) (string, error) {
	creds := credentials.NewStaticCredentials(s.AwsAccessKeyID, s.AwsSecretAccessKey, "")
	cfg := aws.NewConfig().WithRegion(s.AwsS3Region).WithCredentials(creds)
	c := s3.New(session.New(), cfg)
	key := generateKey("png")

	_, err := c.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.AwsS3BucketName),
		Key:         aws.String(key),
		ContentType: aws.String("image/png"),
		Body:        f,
	})

	if err != nil {
		return "", err
	}
	return key, nil
}

func generateKey(ext string) string {
	key := fmt.Sprintf("%s.%s", randHex(8), ext)
	return key
}
