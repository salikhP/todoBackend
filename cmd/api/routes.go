package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.POST("/tasks", app.createTask)
		v1.GET("/tasks", app.getAllTasks)
		v1.GET("/tasks/:id", app.getTask)
		v1.PUT("/tasks/:id", app.updateTask)
		v1.DELETE("events/:id", app.deleteEvent)
	}

	return g
}
