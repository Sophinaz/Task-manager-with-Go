package data

import (
	"context"
	"github.com/Sophinaz/Task-manager-with-Go/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasks() ([]models.Task, error) {
	filter := bson.D{{}}
	var tasks = []models.Task{}
	cur, err := Collection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Task{}, err
	}
	for cur.Next(context.TODO()) {
		var task = models.Task{}
		err = cur.Decode(&task)
		if err != nil {
			log.Fatal("111")
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(id string) (models.Task, error) {
	var task = models.Task{}
	filter := bson.D{{"id", id}}
	err := Collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func UpdateTaskById(updateTask models.Task, id string) error {
	filter := bson.D{{"id", id}}
	update1 := bson.D{{"$set", bson.D{{"title", updateTask.Title}}}}
	update2 := bson.D{{"$set", bson.D{{"description", updateTask.Description}}}}
	_, err := Collection.UpdateOne(context.TODO(), filter, update1)
	_, err = Collection.UpdateOne(context.TODO(), filter, update2)

	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id string) error {
	filter := bson.D{{"id", id}}
	_, err := Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func AddTask(newTask models.Task) error {
	_, err := Collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return err
	}
	return nil
}
