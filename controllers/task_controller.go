package controllers

import (
	"net/http"

	"github.com/Sophinaz/Task-manager-with-Go/data"
	"github.com/Sophinaz/Task-manager-with-Go/models"
	"github.com/gin-gonic/gin"
)


func GetAllTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	c.IndentedJSON(http.StatusOK, task)
}

func UpdateTaskById(c *gin.Context) {
	id := c.Param("id")
	var updateTask models.Task
	err := c.ShouldBindJSON(&updateTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	err = data.UpdateTaskById(updateTask, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func AddTask(c *gin.Context) {
	var newTask models.Task
	err := c.ShouldBindJSON(&newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	err = data.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "task created"})
}
