lint:
	golangci-lint run -c .golangci.yml ./...

test:
	go test -count=1 ./... -covermode=atomic -v -race

build-hello:
	go build -o ./bin/hello ./cmd/hello

build-all: build-hello

.PHONY: lint test
