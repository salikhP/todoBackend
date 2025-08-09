package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoApp/internal/database"
)

func (app *application) createTaskType(c *gin.Context) {
	var taskType database.TaskType

	if err := c.ShouldBindJSON(&taskType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.TaskTypes.Insert(&taskType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, taskType)
}

func (app *application) getAllTaskTypes(c *gin.Context) {
	taskTypes, err := app.models.TaskTypes.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all task types"})
		return
	}

	c.JSON(http.StatusOK, taskTypes)
}

func (app *application) getTaskType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type ID"})
	}

	taskType, err := app.models.TaskTypes.Get(id)

	if taskType == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task type not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task type"})
		return
	}

	c.JSON(http.StatusOK, taskType)
}

func (app *application) updateTaskType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type ID"})
		return
	}

	existingTaskType, err := app.models.TaskTypes.Get(id)

	if existingTaskType == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task type not found"})
		return
	}

	updatedTaskType := &database.TaskType{}

	if err := c.ShouldBindJSON(updatedTaskType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTaskType.Id = id

	if err := app.models.TaskTypes.Update(updatedTaskType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task type"})
	}

	c.JSON(http.StatusOK, updatedTaskType)
}

func (app *application) deleteTaskType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task type ID"})
		return
	}

	if err := app.models.TaskTypes.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task type"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
