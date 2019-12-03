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
		completed bool
		err       int
	}{
		{"title error - small", "a", true, 1},
		{"not completed", "Titulo", false, 0},
		{"title error big ", "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", false, 1},
		{"completed", "Titulo do que fazer", true, 0},
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
