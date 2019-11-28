package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"todo/handlers"

	// "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodo(t *testing.T) {
	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		t.Fatalf("Request was not created %v", err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
