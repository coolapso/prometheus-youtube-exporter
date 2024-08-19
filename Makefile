build:
	go build -o youtube-exporter
fmt:
	go fmt github.com/coolapso/prometheush-youtube-exporter/...
	
build-docker-multiarch:
	docker build --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t youtube-exporter .

test:
	go test ./...
