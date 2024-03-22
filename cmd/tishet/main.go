package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed templates/*
	templates embed.FS
	//go:embed static/*
	static embed.FS
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.Handle("/static/", http.FileServerFS(static))
	log.Fatal(http.ListenAndServe(":8080", Router()))
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	return mux
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		tmpl, err := template.ParseFS(templates, "templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
