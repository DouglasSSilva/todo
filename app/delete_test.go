package app_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/commons"
	"todo/handlers"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	router := handlers.SetupRouter()

	// existingIDs := []int{1, 2, 3, 4}
	nonExistingIDs := []int{1000, 2000, 40000}

	for _, ID := range existingIDs {
		deleted := struct {
			ID uint `json:"ID"`
		}{}

		w := httptest.NewRecorder()
		path := fmt.Sprintf("/api/v1/todos/%d", ID)
		req, err := commons.CreateRequest(nil, "DELETE", path)
		if err != nil {
			t.Fatalf("Failed to created request: %v", err)

		}

		router.ServeHTTP(w, req)

		decoder := json.NewDecoder(w.Body)

		err = decoder.Decode(&deleted)
		if err != nil {
			t.Fatalf("Failed to decode json %v", err)
		}

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, deleted.ID, uint(ID))

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
