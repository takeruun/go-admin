package infrastructure

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsS3 struct {
	Config   *Config
	Uploader *s3manager.Uploader
}

func NewAwsS3() *AwsS3 {
	c := NewConfig()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
				AccessKeyID:     c.AWS.S3.AccessKeyID,
				SecretAccessKey: c.AWS.S3.SecretAccessKey,
			}),
			Endpoint:         aws.String(c.AWS.S3.Endpoint),
			Region:           aws.String(c.AWS.S3.Region),
			S3ForcePathStyle: aws.Bool(true),
		},
	}))

	return &AwsS3{
		Config:   c,
		Uploader: s3manager.NewUploader(sess),
	}
}

func (awsS3 *AwsS3) Upload(file multipart.File, fileName string, extension string) (url string, err error) {
	if fileName == "" {
		return "", errors.New("fileName is required")
	}

	var contentType string

	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "gif":
		contentType = "image/gif"
	case "png":
		contentType = "image/png"
	default:
		return "", errors.New("this extension is invalid")
	}

	result, err := awsS3.Uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Body:        file,
		Bucket:      aws.String(awsS3.Config.AWS.S3.Bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String("images/" + fileName + "." + extension),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}
