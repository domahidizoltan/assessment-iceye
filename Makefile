install:
	go mod vendor
	go install ./...

build:
	go test -cover ./...
	go build -o bin/poker main.go

package: install
	docker build -t iceye/poker .