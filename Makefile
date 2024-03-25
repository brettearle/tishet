templ:
	templ generate

build: templ
	go build -o bin/main cmd/tishet/main.go

run:templ
	go run cmd/tishet/main.go

test:
	go test ./...
