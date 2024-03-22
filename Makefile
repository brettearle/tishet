
build:
	go build -o bin/main cmd/tishet/main.go

run:
	go run cmd/tishet/main.go

test:
	go test -v ./...