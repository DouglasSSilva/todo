package app

import "todo/commons"

// Validate a tTodoModel based on some rules
func (t *TodoModel) Validate() []commons.ErrorMsgs {
	errs := []commons.ErrorMsgs{}

	if t.Title == "" {
		errs = append(errs, commons.ErrorMsgs{
			Field:  "title",
			Motive: "The title field is required",
		})
	}

	if len(t.Title) < 3 || len(t.Title) > 250 {
		errs = append(errs, commons.ErrorMsgs{
			Field:  "title",
			Motive: "The title field must be between 3 - 250 chars",
		})
	}

	return errs
}
