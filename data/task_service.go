package data

import (
	"context"
	"log"

	"go-trial/task-manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasks() ([]models.Task, error) {
	filter := bson.D{{}}
	var tasks = []models.Task{}
	cur, err := Collection.Find(context.TODO(), filter)

	if err != nil {
		return err
	}

	for cur.Next(context.TODO()) {
		var task = models.Task{}
		err = cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
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
