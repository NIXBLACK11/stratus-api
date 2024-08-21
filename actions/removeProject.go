package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"

	"go.mongodb.org/mongo-driver/bson"
)

func RemoveProject(username string, projectname string) (bool, error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return false, errors.New("Database name not present")
	}

	if database.MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	projectsCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Projects")

	err := projectsCollection.FindOneAndDelete(context.TODO(), bson.D{
		{ Key: "username", Value: username },
		{ Key: "projectname", Value: projectname },
	}).Err()
	if err!=nil {
		return false, err
	}

	return true, nil
}