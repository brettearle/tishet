package handlers

import (
	"net/http"

	view "github.com/brettearle/tishet/cmd/tishet/views"
)

func DashBoardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html")
		view := view.DashboardView()
		view.Render(r.Context(), w)
	}
}
