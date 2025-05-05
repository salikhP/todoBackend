package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoApp/internal/database"
)

func (app *application) createTask(c *gin.Context) {
	var task database.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Tasks.insert(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H({"error": "Failed to create task"}))
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (app *application) getAllTasks(c *gin.Context) {
	tasks, err := app.models.Tasks.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all tasks"})
	}

	c.JSON(http.StatusOK, tasks)
}

func (app *application) getTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := app.models.Tasks.Get(id)

	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (app* application) updateTask (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	existingTask, err := app.models.Tasks.Get(id)

	if existingTask == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}

	updatedTask = &database.Task{}

	if err := c.ShouldBindJSON(updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.Id = id

	if err := app.models.Tasks.Update(updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
}
