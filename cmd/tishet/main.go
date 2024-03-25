package main

import (
	"embed"
	"log"
	"net/http"

	handlers "github.com/brettearle/tishet/cmd/tishet/handlers"
)

var (
	//go:embed assets/*
	assets embed.FS
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", Router()))
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.NewHomeHandler(assets))
	mux.HandleFunc("/dashboard", handlers.DashBoardHandler)
	return mux
}
