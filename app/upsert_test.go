package app_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/handlers"

	"github.com/stretchr/testify/assert"
)

func CreateTodoTest(t *testing.T) {

	tt := []struct {
		Name      string `json:"-"`
		Title     string `json:"title"`
		Completed int    `json:"completed"`
		Status    int    `json:"-"`
	}{
		{
			Name:      "working case",
			Title:     "Go to the shopping mall",
			Completed: 0,
			Status:    http.StatusCreated,
		},
	}

	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	fmt.Println("len tc", len(tt))
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest("Post", "/api/v1/todos", nil)
			if err != nil {
				t.Fatalf("Request was not created %v", err)
			}

			router.ServeHTTP(w, req)
			fmt.Println("here")
			assert.Equal(t, tc.Status, w.Code)
		})
	}
}
