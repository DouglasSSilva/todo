package main

import (
	"todo/handlers"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := handlers.SetupRouter()
	r.Run(":8080")
}
