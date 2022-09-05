# DownTube

Go Application that Converts YouTube videos to MP3s

## Tech

This service has 2 main dependencies
* [goutubedl](https://github.com/wader/goutubedl) library to download the YouTube video
* [fluent-ffmpeg](https://github.com/modfy/go-fluent-ffmpeg) to convert the video to MP3

## Local Instructions

* Ensure [docker](https://www.docker.com/) is installed
* Build the docker image
```
docker build -t latest .
```
* To run locally:
```
docker run -p 8080:8080 downtube:latest
```