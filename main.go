package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"github.com/wader/goutubedl"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/download", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	vID := r.URL.Query().Get("videoId")
	if vID == "" {
		log.Fatal("Empty video ID")
	}

	dBuff := &bytes.Buffer{}
	goutubedl.Path = "yt-dlp"
	result, err := goutubedl.New(context.Background(), buildYtURL(vID), goutubedl.Options{})
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(dBuff, downloadResult)
	if err != nil {
		log.Fatal(err)
	}

	mBuff := &bytes.Buffer{}
	err = fluentffmpeg.
		NewCommand("").
		PipeInput(dBuff).
		OutputFormat("mp3").
		PipeOutput(mBuff).
		Run()
	if err != nil {
		log.Fatal(err)
	}

	w.Write(mBuff.Bytes())
}

func buildYtURL(videoID string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
}
