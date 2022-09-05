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

	result, err := goutubedl.New(context.Background(), buildYTURL(videoID), goutubedl.Options{})
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

func buildYTURL(videoID string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
}
