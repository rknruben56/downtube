// Package download handles the logic for downloading videos
package download

//go:generate mockgen -destination=mock/download.go . Downloader

// Downloader handles the video interaction
type Downloader interface {
	Download(videoID string) (Result, error)
}
