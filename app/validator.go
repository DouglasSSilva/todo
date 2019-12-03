package app

import "todo/commons"

// Validate a tTodoModel based on some rules
func (t *TodoModel) Validate() []commons.ErrorToReturn {
	errs := []commons.ErrorToReturn{}

	if t.Title == "" {
		errs = append(errs, commons.ErrorToReturn{
			Field:  "title",
			Motive: "The title field is required",
		})
	}

	if len(t.Title) < 3 || len(t.Title) > 250 {
		errs = append(errs, commons.ErrorToReturn{
			Field:  "title",
			Motive: "The title field must be between 3 - 250 chars",
		})
	}

	if t.Completed < 0 || t.Completed > 1 {
		errs = append(errs, commons.ErrorToReturn{
			Field:  "completed",
			Motive: "The completed value can not be recognized",
		})
	}

	return errs
}
