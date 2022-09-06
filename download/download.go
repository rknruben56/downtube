package download

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/wader/goutubedl"
)

const downloadQuality = "best"

// YTDownloader downloads the video from YouTube
type YTDownloader struct {
	Path string
}

// Download downloads the video from YouTube and returns a populated Buffer if there
// are no errors
func (d *YTDownloader) Download(videoID string) (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	goutubedl.Path = d.Path

	result, err := d.GetInfo(videoID)
	if err != nil {
		return b, nil
	}

	downloadResult, err := result.Download(context.Background(), downloadQuality)
	if err != nil {
		return b, nil
	}

	_, err = io.Copy(b, downloadResult)
	return b, err
}

// GetInfo returns the metadata of the YouTube Video
func (d *YTDownloader) GetInfo(videoID string) (goutubedl.Result, error) {
	return goutubedl.New(context.Background(), buildYTURL(videoID), goutubedl.Options{})
}

func buildYTURL(videoID string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
}
