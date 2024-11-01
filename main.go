package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"duedate"`
	Status      string    `json:"status"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal((err))
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("mongo connected")

	collection := client.Database("Task_manager").Collection("tasks")

	router := gin.Default()
	router.GET("/tasks", func(ctx *gin.Context) { getAllTasks(ctx, collection) })
	router.GET("/tasks/:id", func(ctx *gin.Context) { getTaskById(ctx, collection) })
	router.PUT("/tasks/:id", func(ctx *gin.Context) {updateTaskById(ctx, collection)} )
	router.DELETE("tasks/:id", func(ctx *gin.Context) {deleteTask(ctx, collection)})
	router.POST("/tasks", func(ctx *gin.Context) { addTask(ctx, collection) })

	router.Run()
}

func getAllTasks(c *gin.Context, collection *mongo.Collection) {
	fmt.Println(collection)
	filter := bson.D{{}}
	var tasks = []Task{}
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var task = Task{}
		err = cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskById(c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")
	var task = Task{}

	filter := bson.D{{"id", id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	c.IndentedJSON(http.StatusOK, task)
}

func updateTaskById(c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")

	var updateTask Task

	err := c.ShouldBindJSON(&updateTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	filter := bson.D{{"id", id}}
	update1 := bson.D{{"$set", bson.D{{"title", updateTask.Title}}}}
	update2 := bson.D{{"$set", bson.D{{"description", updateTask.Description}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update1)
	_, err = collection.UpdateOne(context.TODO(), filter, update2)

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func deleteTask(c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")

	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted"})

}

func addTask(c *gin.Context, collection *mongo.Collection) {

	var newTask Task

	err := c.ShouldBindJSON(&newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(newTask)

	_, err = collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task created"})
}
