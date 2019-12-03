package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []TodoModel

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
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
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
