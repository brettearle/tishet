package main

import (
	"testing"
)

func TestRouter(t *testing.T) {
	t.Run("Router returns a ServeMux", func(t *testing.T) {
		mux := Router()
		if mux == nil {
			t.Errorf("Expected a ServeMux but got nil")
		}
	})
}
