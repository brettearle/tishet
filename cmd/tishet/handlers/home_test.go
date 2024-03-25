package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"
)

func TestHomeHandler(t *testing.T) {
	mockFile := fstest.MapFile{
		Data: []byte("<html lang=\"en\"><head><title>Home</title></head><body></body></html>"),
	}
	mockAssets := fstest.MapFS{}
	mockAssets["index.html"] = &mockFile

	homeHandler := NewHomeHandler(mockAssets)
	t.Run(" '/' route returns status 200", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()
		homeHandler(response, request)
		if response.Result().StatusCode != 200 {
			t.Errorf("Expected Status %v but got %v", 200, response.Result().StatusCode)
		}
	})
	t.Run(" '/' route returns html content type", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()
		homeHandler(response, request)
		contentHeader := response.Header().Get("Content-Type")
		if contentHeader != "text/html" {
			t.Errorf("Expected Content-Type %v but got %v", "text/html", contentHeader)
		}
	})

	t.Run(" '/' returns a html body", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/", nil)
		response := httptest.NewRecorder()
		homeHandler(response, request)
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
