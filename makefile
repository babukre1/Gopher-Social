build:
	@go build -o bin/com cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/com 