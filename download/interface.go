// Package download handles the logic for downloading videos
package download

import (
	"bytes"

	"github.com/wader/goutubedl"
)

// Downloader handles the video interaction
type Downloader interface {
	Download(videoID string) (*bytes.Buffer, error)
	GetInfo(videoID string) (goutubedl.Result, error)
}
