package main

import (
	"embed"
	"log"
	"net/http"

	view "github.com/brettearle/tishet/cmd/tishet/views"
)

var (
	//go:embed assets/*
	assets embed.FS
)

func main() {
	http.Handle("/assets/", http.FileServerFS(assets))
	log.Fatal(http.ListenAndServe(":8080", Router()))
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/dashboard", DashBoardHandler)
	return mux
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		view := view.HomeView()
		view.Render(r.Context(), w)
	}
}

func DashBoardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		view := view.DashboardView()
		view.Render(r.Context(), w)
	}
}
