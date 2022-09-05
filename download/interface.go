// Package download handles the logic for downloading videos
package download

import "bytes"

// Downloader downloads the video
type Downloader interface {
	Download(videoID string) (*bytes.Buffer, error)
}
