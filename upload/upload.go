package upload

import (
	"bytes"
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AWSUploader uploads to AWS
type AWSUploader struct {
	Bucket string
	Config aws.Config
}

const contentType = "audio/mpeg"
const keyPrefix = "songs/"
const timeout = 15

// Upload uploads the MP3 buffer to S3 and retrieves a signed URL
func (u *AWSUploader) Upload(title string, buffer *bytes.Buffer) (string, error) {
	sess, err := session.NewSession(&u.Config)
	if err != nil {
		return "", nil
	}

	key := keyPrefix + title
	s3Uploader := s3manager.NewUploader(sess)
	input := &s3manager.UploadInput{
		Bucket:      aws.String(u.Bucket),
		Key:         aws.String(key),
		Body:        buffer,
		ContentType: aws.String(contentType),
	}

	_, err = s3Uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(timeout * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}
