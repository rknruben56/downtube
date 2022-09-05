// Package main starts the web server
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rknruben56/downtube/download"
	"github.com/rknruben56/downtube/transcode"
)

var downloader download.Downloader
var transcoder transcode.Transcoder

func main() {
	log.Print("starting server...")
	initComponents()
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
		err := fmt.Errorf("Invalid video ID: %s", vID)
		handleError(w, http.StatusBadRequest, err)
		return
	}

	dBuff, err := downloader.Download(vID)
	if err != nil {
		err = fmt.Errorf("Error downloading video: %s", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	tBuff, err := transcoder.Transcode(dBuff)
	if err != nil {
		err = fmt.Errorf("Error transcoding video: %s", err)
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(tBuff.Bytes())
}

func initComponents() {
	downloader = &download.YTDownloader{Path: "yt-dlp"}
	transcoder = &transcode.MP3Transcoder{}
}

func handleError(w http.ResponseWriter, status int, err error) {
	log.Println(err)
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}
