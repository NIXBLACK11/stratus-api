package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"
	"stratus-api/models"
)

func CreateUser(user models.User) (bool, error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return false, errors.New("Database name not present")
	}

	if database.MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	usersCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Users")

	_, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		return false, err
	}

	return true, nil
}