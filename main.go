package main

import (
	"github.com/Sophinaz/Task-manager-with-Go/controllers"
	"github.com/Sophinaz/Task-manager-with-Go/data"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.PUT("/tasks/:id", controllers.UpdateTaskById)
	router.DELETE("tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.AddTask)
	data.Connect()
	router.Run()
}
