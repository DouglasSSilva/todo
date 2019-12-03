package app_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"todo/app"
	"todo/commons"
	"todo/handlers"

	"github.com/stretchr/testify/assert"
)

//TestCreateTodo based on a json as a table test.
// checks for the return code and related structure.
func TestCreateTodo(t *testing.T) {

	tt := []struct {
		Name        string `json:"name"`
		Title       string `json:"title"`
		Completed   int    `json:"completed"`
		Status      int    `json:"status"`
		TotalErrors int    `json:"totalErrors"`
	}{}

	err := commons.GetJSONTestFiles(&tt, "testfiles/insert_cases.json")
	if err != nil {
		t.Fatalf("Failed to read json %v", err)
	}

	router := handlers.SetupRouter()

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, err := commons.CreateRequest(tc, "POST", "/api/v1/todos")
			if err != nil {
				t.Fatalf("Failed to created request: %v", err)
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.Status, w.Code)
			fmt.Println(w.Body)

			decoder := json.NewDecoder(w.Body)
			if w.Code == 201 {

				//evaluate the creation of a todo
				todo := app.TodoModel{}

				err := decoder.Decode(&todo)
				if err != nil {
					t.Fatalf("Failed to decode json %v", err)
				}
				assert.Equal(t, tc.Title, todo.Title)
				assert.Equal(t, tc.Completed, todo.Completed)
				assert.Greater(t, todo.ID, uint(0))
			} else {

				//evaluate when there is an error over the todo creation
				errs := []commons.ErrorToReturn{}
				err := decoder.Decode(&errs)
				if err != nil {
					t.Fatalf("Failed to decode json %v", err)
				}

				assert.Equal(t, len(errs), tc.TotalErrors)
			}
		})
	}
}
