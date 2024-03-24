
build:
	go build -o bin/main cmd/tishet/main.go

templ:
	templ generate

run:templ
	go run cmd/tishet/main.go

test:
	go test ./...
