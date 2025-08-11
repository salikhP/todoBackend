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
		v1.DELETE("tasks/:id", app.deleteTask)

		v1.POST("/taskTypes", app.createTaskType)
		v1.GET("/taskTypes", app.getAllTaskTypes)
		v1.GET("/taskTypes/:id", app.getTaskType)
		v1.PUT("/taskTypes/:id", app.updateTaskType)
		v1.DELETE("/taskTypes/:id", app.deleteTaskType)

		v1.GET("/users/:userId/tasks", app.listTasks)

		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.login)
	}

	return g
}
