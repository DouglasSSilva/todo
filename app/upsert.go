package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodo add a new todo
func CreateTodo(c *gin.Context) {
	todo := TodoModel{}
	c.BindJSON(&todo)

	err := todo.Validate()
	if len(err) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Error": err})
		return
	}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// UpdateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var todo TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}
