package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodo add a new todo
func CreateTodo() gin.HandlerFunc {
	fmt.Println("here")
	return func(c *gin.Context) {
		completed, _ := strconv.Atoi(c.PostForm("completed"))
		todo := TodoModel{Title: c.PostForm("title"), Completed: completed}

		db.Save(&todo)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
	}
}

// FetchAllTodo fetch all todos
func FetchAllTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []TodoModel
		var _todos []TransformedTodo

		db.Find(&todos)

		if len(todos) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
			return
		}

		//transforms the todos for building a good response
		for _, item := range todos {
			completed := false
			if item.Completed == 1 {
				completed = true
			} else {
				completed = false
			}
			_todos = append(_todos, TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
	}
}

// FetchSingleTodo fetch a single todo
func FetchSingleTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo TodoModel
		todoID := c.Param("id")

		db.First(&todo, todoID)

		if todo.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
			return
		}

		completed := false
		if todo.Completed == 1 {
			completed = true
		} else {
			completed = false
		}

		_todo := TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
	}
}

// UpdateTodo update a todo
func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// DeleteTodo remove a todo
func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo TodoModel
		todoID := c.Param("id")

		db.First(&todo, todoID)

		if todo.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
			return
		}

		db.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
	}
}
