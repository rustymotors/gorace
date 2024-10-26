test:
	go test -count=1  -coverprofile=coverage.txt -covermode=atomic ./internal/...

build:
	go build -o bin/server cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf bin

run_bin:
	./bin/server

.PHONY: test build run clean run_bin