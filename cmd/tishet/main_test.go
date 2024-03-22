package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("Hello, World!")
	t.Run("Test Hello World", func(t *testing.T) {
		t.Log("Hello, from t run")
	})
}

