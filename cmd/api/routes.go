package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")

	// public routes
	{
		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.login)
	}

	// protected routes (with AuthMiddleware)
	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleware())
	{
		authGroup.POST("/tasks", app.createTask)
		authGroup.GET("/tasks", app.getAllTasks)
		authGroup.GET("/tasks/:id", app.getTask)
		authGroup.PUT("/tasks/:id", app.updateTask)
		authGroup.DELETE("tasks/:id", app.deleteTask)

		authGroup.POST("/taskTypes", app.createTaskType)
		authGroup.GET("/taskTypes", app.getAllTaskTypes)
		authGroup.GET("/taskTypes/:id", app.getTaskType)
		authGroup.PUT("/taskTypes/:id", app.updateTaskType)
		authGroup.DELETE("/taskTypes/:id", app.deleteTaskType)

		authGroup.GET("/users/:userId/tasks", app.listTasks)
	}

	return g
}
