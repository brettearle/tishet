package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	http.ListenAndServe(":8080", Router())
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	return mux
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	indexTmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = indexTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Home</title>
</head>
<body>
	<h1>Welcome to Tishet</h1>
</body>
</html>
`