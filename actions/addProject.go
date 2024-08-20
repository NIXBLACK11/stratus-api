package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"
	"stratus-api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func AddProject(username string, project models.Project) (bool ,error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return false, errors.New("Database name not present")
	}

	if database.MongoDBClient == nil {
		return false, errors.New("MongoDB client is not initialized")
	}

	if userExists, err := CheckUserExists(username); userExists==false {
		return false, err
	}

	if username!=project.UserName {
		return false, errors.New("Invalid request")
	}

	projectsCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Projects")

	var projectDelete bson.M
	_ = projectsCollection.FindOneAndDelete(context.TODO(), bson.D{
		{ Key: "username", Value: username},
		{ Key: "projectname", Value: project.ProjectName},
	}).Decode(&projectDelete)

	_, err := projectsCollection.InsertOne(context.Background(), project)
	if err!=nil {
		return false, err
	}

	return true, nil
}