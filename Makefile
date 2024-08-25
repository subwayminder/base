GOBASE=$(shell pwd)

compile:
	@echo "  >  Building binaries..."
	GOOS=linux GOARCH=amd64 go build -o bin/base
	GOOS=linux GOARCH=arm64 go build -o bin/base-linux-arm
	GOOS=windows GOARCH=amd64 go build -o bin/base.exe
	GOOS=darwin GOARCH=amd64 go build -o bin/base-amd64
	GOOS=darwin GOARCH=arm64 go build -o bin/base-arm64
	@echo "  >  Done."