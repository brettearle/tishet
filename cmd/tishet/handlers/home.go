package handlers

import (
	"io/fs"
	"net/http"
	"strings"

	views "github.com/brettearle/tishet/cmd/tishet/views"
)

func NewHomeHandler(assets fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//serve css and js files for home route
		if strings.Contains(r.URL.Path, "assets") {
			http.FileServer(http.FS(assets)).ServeHTTP(w, r)
			return
		}

		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "text/html")
			view := views.HomeView()
			view.Render(r.Context(), w)
		}
	}
}
