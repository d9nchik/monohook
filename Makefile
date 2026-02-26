.PHONY: build

build:
	rm -rf build
	mkdir -p  build
	make build-monohook


build-monohook:
	cd build && \
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o bootstrap ../cmd/main.go && \
	zip monohook.zip bootstrap && \
	rm bootstrap