package app_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/app"
	"todo/commons"
	"todo/handlers"

	"github.com/stretchr/testify/assert"
)

// "net/http/httptest"
func TestFetchAll(t *testing.T) {

	router := handlers.SetupRouter()

	w := httptest.NewRecorder()
	path := "/api/v1/todos"
	req, err := commons.CreateRequest(nil, "GET", path)
	if err != nil {
		t.Fatalf("Failed to created request: %v", err)
	}
	router.ServeHTTP(w, req)

	decoder := json.NewDecoder(w.Body)
	todo := []app.TodoModel{}
	err = decoder.Decode(&todo)
	if err != nil {
		errs := []commons.ErrorMsgs{}
		err = decoder.Decode(&errs)
		if err != nil {
			t.Fatalf("Failed to decode JSON %+v", err)

		}
		assert.Equal(t, w.Code, http.StatusNotFound)
		assert.Equal(t, errs[0].Field, "Todo")
		assert.Equal(t, errs[0].Motive, "Not Found")
		assert.Equal(t, len(todo), 0)
	} else {
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Greater(t, len(todo), 0)
	}
}

func TestFetchByID(t *testing.T) {

	router := handlers.SetupRouter()

	existingIDs := []int{1, 2, 3, 4}
	nonExistingIDs := []int{1000, 2000, 40000}

	for _, ID := range existingIDs {
		w := httptest.NewRecorder()
		path := fmt.Sprintf("/api/v1/todos/%d", ID)
		req, err := commons.CreateRequest(nil, "GET", path)
		if err != nil {
			t.Fatalf("Failed to created request: %v", err)

		}

		router.ServeHTTP(w, req)

		decoder := json.NewDecoder(w.Body)
		todo := app.TodoModel{}

		err = decoder.Decode(&todo)
		if err != nil {
			t.Fatalf("Failed to decode json %v", err)
		}

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, todo.ID, uint(ID))
		assert.Greater(t, len(todo.Title), 3)
		assert.Less(t, len(todo.Title), 250)

	}

	for _, ID := range nonExistingIDs {
		w := httptest.NewRecorder()
		path := fmt.Sprintf("/api/v1/todos/%d", ID)
		req, err := commons.CreateRequest(t, "GET", path)
		if err != nil {
			t.Fatalf("Failed to created request: %v", err)
		}

		router.ServeHTTP(w, req)

		decoder := json.NewDecoder(w.Body)
		errs := []commons.ErrorMsgs{}
		err = decoder.Decode(&errs)
		if err != nil {
			t.Fatalf("Failed to decode json %v", err)
		}

		assert.Equal(t, w.Code, http.StatusNotFound)
		assert.Equal(t, errs[0].Field, "Todo")
		assert.Equal(t, errs[0].Motive, "Not Found")
	}
}
