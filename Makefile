.PHONY: build

build:
	rm -rf build
	mkdir -p  build
	make build-monohook


build-monohook:
	cd build && \
	GOARCH=amd64 GOOS=linux go build -o bootstrap ../cmd/main.go && \
	zip monohook.zip bootstrap && \
	rm bootstrap