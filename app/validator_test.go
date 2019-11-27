package app_test

import (
	"fmt"
	"testing"
	"todo/app"
)

func TestValidator(t *testing.T) {
	tModel := app.TodoModel{}

	tt := []struct {
		name      string
		title     string
		completed int
		err       int
	}{
		{"title error", "a", 1, 1},
		{"not completed", "Titulo", 0, 0},
		{"completed error", "Titulo", 2, 1},
		{"title and completed error", "ab", 2, 2},
		{"completed", "Titulo do que fazer", 1, 0},
	}

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tModel.Title = tc.title
			tModel.Completed = tc.completed
			errs := tModel.Validate()
			if len(errs) != tc.err {
				err := fmt.Sprintf("Test %d  failed with %d errors", i, len(errs))
				t.Fatalf(err)
			}
		})
	}

}
