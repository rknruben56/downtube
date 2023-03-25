// Package upload interacts with AWS
package upload

import "bytes"

//go:generate mockgen -destination=mock/upload.go . Uploader

// Uploader uploads the MP3 to S3 and returns a signed URL
type Uploader interface {
	Upload(title string, buffer *bytes.Buffer) (string, error)
}
