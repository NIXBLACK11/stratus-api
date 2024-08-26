package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserExists(username string) (bool, error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return false, errors.New("Database name not present")
	}

	if database.MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	usersCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Users")

	var user bson.M
	err := usersCollection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
	}).Decode(&user)
	if err!=nil {
		if err==mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}