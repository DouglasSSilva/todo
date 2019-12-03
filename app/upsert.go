package app

import (
	"fmt"
	"net/http"
	"todo/commons"

	"github.com/gin-gonic/gin"
)

// CreateTodo add a new todo
func CreateTodo(c *gin.Context) {
	todo := TodoModel{}
	c.BindJSON(&todo)

	err := todo.Validate()
	if len(err) != 0 {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	db.Save(&todo)
	c.JSON(http.StatusCreated, todo)

}

// UpdateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var todo, updateTodo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		errs := []commons.ErrorMsgs{}
		errs = append(errs, commons.ErrorMsgs{Field: "Todo", Motive: "No such id in the database"})
		c.JSON(http.StatusNotFound, errs)
		return
	}

	c.BindJSON(&updateTodo)
	err := updateTodo.Validate()
	if len(err) != 0 {
		fmt.Println("here")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	db.Save(&updateTodo)
	c.JSON(http.StatusOK, updateTodo)
}
