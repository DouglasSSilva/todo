package app_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/handlers"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	tt := []struct {
		Name        string `json:"-"`
		Title       string `json:"title"`
		Completed   int    `json:"completed"`
		Status      int    `json:"-"`
		TotalErrors int    `json:"-"`
	}{
		{
			Name:        "working finished case",
			Title:       "Todo create test",
			Completed:   1,
			Status:      http.StatusCreated,
			TotalErrors: 0,
		},
		{
			Name:        "working unfinished case",
			Title:       "Todo update test",
			Completed:   0,
			Status:      http.StatusCreated,
			TotalErrors: 0,
		},
		{
			Name:        "chars error",
			Title:       "Go",
			Completed:   0,
			Status:      http.StatusInternalServerError,
			TotalErrors: 1,
		},
		{
			Name:        "completed error",
			Title:       "Go and rest",
			Completed:   2,
			Status:      http.StatusInternalServerError,
			TotalErrors: 1,
		},
		{
			Name:        "all - chars and completed error",
			Title:       "Go",
			Completed:   2,
			Status:      http.StatusInternalServerError,
			TotalErrors: 2,
		},
	}

	router := handlers.SetupRouter()

	fmt.Println("len tc", len(tt))
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b, err := json.Marshal(tc)
			if err != nil {
				t.Fatalf("Body was not created %v", err)
			}
			bodyBuffer := bytes.NewBuffer(b)
			fmt.Println(bodyBuffer)
			req, err := http.NewRequest("POST", "/api/v1/todos", bodyBuffer)
			if err != nil {
				t.Fatalf("Request was not created %v", err)
			}

			router.ServeHTTP(w, req)
			assert.Equal(t, tc.Status, w.Code)
		})
	}
}
