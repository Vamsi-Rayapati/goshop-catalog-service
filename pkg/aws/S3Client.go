package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/smartbot/catalog/pkg/config"
	"github.com/smartbot/catalog/pkg/errors"
)

func GetS3Client() *s3.Client {
	creds := credentials.NewStaticCredentialsProvider(config.Config.AWSApiKey, config.Config.AWSSecret, "")
	config := aws.Config{
		Region:      "us-east-2",
		Credentials: creds,
	}
	

	client := s3.NewFromConfig(config)
	return client
}
func GetNewSignedUrl(bucketName string, fileName string, contentType string) (*string, *string, *errors.ApiError) {
	s3Client := GetS3Client()
	presigner := s3.NewPresignClient(s3Client)
	key := uuid.NewString() + "_" + fileName
	params := &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &key,
		ContentType: &contentType,
	}

	presignedReq, err := presigner.PresignPutObject(context.TODO(), params, s3.WithPresignExpires(5*time.Minute))
	if err != nil {
		return nil, nil, errors.InternalServerError("Failed to generate upload url")
	}

	return &presignedReq.URL, &key, nil
}

func GetSignedUrl(bucketName string, key string) (*string, *errors.ApiError) {
	s3Client := GetS3Client()
	presigner := s3.NewPresignClient(s3Client)

	params := &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	}

	presignedReq, err := presigner.PresignGetObject(context.TODO(), params, s3.WithPresignExpires(5*time.Minute))
	if err != nil {
		return nil, errors.InternalServerError("Failed to generate upload url")

	}

	return &presignedReq.URL, nil
}
