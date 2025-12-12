package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskModel struct {
	ID          int    `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	DueDate     string `json:"dueDate" bson:"dueDate"`
	Status      bool   `json:"status" bson:"status"`
}

var taskCollection *mongo.Collection

func init() {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}

	log.Printf("MongoDB connection URL: %s", mongoURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
	}

	db := client.Database("task_manager_db")
	taskCollection = db.Collection("tasks")
}

func getNextTaskID(ctx context.Context) (int, error) {
	var last TaskModel
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := taskCollection.FindOne(ctx, bson.D{}, opts).Decode(&last)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return 1, nil
	}
	if err != nil {
		return 0, err
	}
	return last.ID + 1, nil
}

func GetAllTasks() []TaskModel {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("error retrieving tasks from MongoDB: %v", err)
		return []TaskModel{}
	}
	defer cursor.Close(ctx)

	var tasks []TaskModel
	for cursor.Next(ctx) {
		var task TaskModel
		if err := cursor.Decode(&task); err != nil {
			log.Printf("error decoding task document: %v", err)
			continue
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("cursor error while reading tasks: %v", err)
	}

	return tasks
}

func GetTaskByID(id int) (TaskModel, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task TaskModel
	err := taskCollection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return TaskModel{}, false
	}
	if err != nil {
		log.Printf("error fetching task with id=%d: %v", id, err)
		return TaskModel{}, false
	}

	return task, true
}

func CreateTask(newTask TaskModel) TaskModel {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nextID, err := getNextTaskID(ctx)
	if err != nil {
		log.Printf("error generating next task ID: %v", err)
		return TaskModel{}
	}
	newTask.ID = nextID

	_, err = taskCollection.InsertOne(ctx, newTask)
	if err != nil {
		log.Printf("error inserting task into MongoDB: %v", err)
		return TaskModel{}
	}

	return newTask
}

func UpdateTask(id int, updatedDetails TaskModel) (TaskModel, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updatedDetails.ID = id

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedDetails}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated TaskModel
	err := taskCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return TaskModel{}, false
	}
	if err != nil {
		log.Printf("error updating task with id=%d: %v", id, err)
		return TaskModel{}, false
	}

	return updated, true
}

func DeleteTask(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := taskCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Printf("error deleting task with id=%d: %v", id, err)
		return false
	}

	return result.DeletedCount > 0
}

