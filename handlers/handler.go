package handlers

import (
	"todo/app"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("", app.CreateTodo)
		v1.GET("", app.FetchAllTodo)
		v1.GET("/:id", app.FetchSingleTodo)
		v1.PUT("/:id", app.UpdateTodo)
		v1.DELETE("/:id", app.DeleteTodo)
	}

	return router
}
