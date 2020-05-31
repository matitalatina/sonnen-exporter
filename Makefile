.PHONY: build build-arm

BUILD_CMD=go build -o dist/sonnen-exporter cmd/sonnen-exporter/main.go

build-arm:
	GOOS=linux GOARCH=arm GOARM=7 $(BUILD_CMD)

build:
	$(BUILD_CMD)
