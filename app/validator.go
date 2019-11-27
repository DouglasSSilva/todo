package app

import "net/url"

// Validate a tTodoModel based on some rules
func (t *TodoModel) Validate() url.Values {
	errs := url.Values{}

	if t.Title == "" {
		errs.Add("title", "The title field is required")
	}

	if len(t.Title) < 3 || len(t.Title) > 250 {
		errs.Add("title", "The title field must be between 3 - 250 chars")
	}

	if t.Completed < 0 || t.Completed > 1 {
		errs.Add("completed", "The completed value can not be recognized")
	}

	return errs
}
