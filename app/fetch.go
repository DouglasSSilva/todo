package app

import (
	"fmt"
	"net/http"
	"todo/commons"

	"github.com/gin-gonic/gin"
)

// FetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []TodoModel

	db.Find(&todos)

	if len(todos) <= 0 {
		errs := []commons.ErrorMsgs{}
		errs = append(errs, commons.ErrorMsgs{
			Field:  "Todo",
			Motive: "Not Found"})
		fmt.Println(errs)
		c.JSON(http.StatusNotFound, errs)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// FetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		errs := []commons.ErrorMsgs{}
		errs = append(errs, commons.ErrorMsgs{
			Field:  "Todo",
			Motive: "Not Found"})

		c.JSON(http.StatusNotFound, errs)
		return
	}

	c.JSON(http.StatusOK, todo)
}
