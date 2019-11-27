package main

import (
	"net/http"
	"net/http/httptest"
	"todo/app"

	// "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodo(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "localhost:8080/api/v1/todos/", nil)
	if err != nil {
		t.Fatalf("Request was not created %v", err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	app.FetchAllTodo()
}
