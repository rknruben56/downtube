ARG GOLANG_VERSION=1.19.0
ARG YT_DLP=2023.12.30
ARG FFMPEG_VERSION=5.1.1

# Build App
FROM golang:${GOLANG_VERSION} as builder
WORKDIR /app
COPY go.* *.go ./
RUN go mod download
COPY . ./
RUN go build -v -o server

# Setup yt-dlp
FROM golang:$GOLANG_VERSION AS yt-dlp
ARG YT_DLP
RUN \
  curl -L https://github.com/yt-dlp/yt-dlp/releases/download/${YT_DLP}/yt-dlp -o /yt-dlp && \
  chmod a+x /yt-dlp

# Setup ffmpeg
FROM mwader/static-ffmpeg:$FFMPEG_VERSION AS ffmpeg

# Setup Image
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y -q \
    ca-certificates python3 curl && \
    rm -rf /var/lib/apt/lists/*
RUN ln -s /usr/bin/python3 /usr/bin/python

COPY --from=yt-dlp /yt-dlp /usr/local/bin/
COPY --from=ffmpeg /ffmpeg /ffprobe /usr/local/bin/
COPY --from=builder /app/server /app/server

RUN yt-dlp --version && \
    ffmpeg -version && \
    ffprobe -version

CMD ["/app/server"]