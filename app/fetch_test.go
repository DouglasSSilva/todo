package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"todo/app"
	"todo/handlers"

	// "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	Data []app.TransformedTodo `json:"data"`
}

func TestFetchAllTodo(t *testing.T) {
	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/todos", nil)
	if err != nil {
		t.Fatalf("Request was not created %v", err)
	}

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	resp := Response{}
	err = json.Unmarshal([]byte(w.Body.String()), &resp)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(resp.Data), 1)

}
