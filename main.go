package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"os"

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
	dBuff := &bytes.Buffer{}
	goutubedl.Path = "yt-dlp"
	result, err := goutubedl.New(context.Background(), "https://www.youtube.com/watch?v=vSffKUyr0lk", goutubedl.Options{})
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(dBuff, downloadResult)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(dBuff.Bytes())
}
