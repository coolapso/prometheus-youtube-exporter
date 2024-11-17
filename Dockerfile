FROM --platform=$BUILDPLATFORM golang:latest AS builder
ARG TARGETARCH

WORKDIR /youtube-exporter
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags="-X github.com/coolapso/go-live-server/cmd.Version=${VERSION}" -a -o youtube-exporter

FROM alpine:latest

COPY --from=builder youtube-exporter/youtube-exporter /usr/bin/youtube-exporter

EXPOSE 10020
ENTRYPOINT ["/usr/bin/youtube-exporter"]
