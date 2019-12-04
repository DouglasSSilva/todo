package app

import (
	"net/http"
	"todo/commons"

	"github.com/gin-gonic/gin"
)

// DeleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
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

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"ID": todo.ID})
}
