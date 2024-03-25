package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDashBoardHandler(t *testing.T) {
	t.Run(" '/dashboard' route returns status 200", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/dashboard", nil)
		response := httptest.NewRecorder()
		DashBoardHandler(response, request)
		if response.Result().StatusCode != 200 {
			t.Errorf("Expected Status %v but got %v", 200, response.Result().StatusCode)
		}
	})
	t.Run(" '/dashboard' route returns html content type", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()
		DashBoardHandler(response, request)
		contentHeader := response.Header().Get("Content-Type")
		if contentHeader != "text/html" {
			t.Errorf("Expected Content-Type %v but got %v", "text/html", contentHeader)
		}
	})
	t.Run(" '/dashboard' returns a html body", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/dashboard", nil)
		response := httptest.NewRecorder()
		DashBoardHandler(response, request)
		body := response.Body.String()
		//test if the body is a valid html document without checking the content
		if !strings.Contains(body, "<html lang=\"en\"") || !strings.Contains(body, "</html>") {
			t.Errorf("Expected Body to be a valid html document")
		}
		//contains a title
		if !strings.Contains(body, "<title>") || !strings.Contains(body, "</title>") {
			t.Errorf("Expected Body to contain a title")
		}
	})
}
