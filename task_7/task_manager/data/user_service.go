package database

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"-" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

var (
	userCollection *mongo.Collection
	userInitOnce   sync.Once
)

func initUsers() {
	userInitOnce.Do(func() {
		mongoURL := os.Getenv("MONGO_URL")
		if mongoURL == "" {
			mongoURL = "mongodb://localhost:27017"
		}

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
		userCollection = db.Collection("users")
	})
}

func getNextUserID(ctx context.Context) (int, error) {
	var last UserModel
	opts := options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})
	err := userCollection.FindOne(ctx, bson.D{}, opts).Decode(&last)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return 1, nil
	}
	if err != nil {
		return 0, err
	}
	return last.ID + 1, nil
}

func isDatabaseEmpty(ctx context.Context) bool {
	count, err := userCollection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Printf("error counting users: %v", err)
		return false
	}
	return count == 0
}

func CreateUser(username, password string) (UserModel, error) {
	initUsers()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser, _ := GetUserByUsername(username)
	if existingUser.Username != "" {
		return UserModel{}, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return UserModel{}, err
	}

	nextID, err := getNextUserID(ctx)
	if err != nil {
		return UserModel{}, err
	}

	role := "user"
	if isDatabaseEmpty(ctx) {
		role = "admin"
	}

	user := UserModel{
		ID:       nextID,
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return UserModel{}, err
	}

	user.Password = ""
	return user, nil
}

func GetUserByUsername(username string) (UserModel, error) {
	initUsers()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user UserModel
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return UserModel{}, errors.New("user not found")
	}
	if err != nil {
		return UserModel{}, err
	}

	return user, nil
}

func GetUserByID(id int) (UserModel, error) {
	initUsers()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user UserModel
	err := userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return UserModel{}, errors.New("user not found")
	}
	if err != nil {
		return UserModel{}, err
	}

	return user, nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func PromoteUser(username string) error {
	initUsers()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

