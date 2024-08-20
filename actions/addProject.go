package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"
	"stratus-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	usersCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Users")

	var user bson.M
	err := usersCollection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
	}).Decode(&user)
	if err!=nil {
		return false, err
	}

	projectname := project.ProjectName
	userId, ok := user["_id"].(primitive.ObjectID)
	if !ok {
		return false, nil
	}

	projectsCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Projects")

	var projectDelete bson.M
	err = projectsCollection.FindOneAndDelete(context.TODO(), bson.D{
		{ Key: "projectid", Value: userId},
		{ Key: "projectname", Value: projectname},
	}).Decode(&projectDelete)

	project.UserId = userId
	_, err = projectsCollection.InsertOne(context.Background(), project)
	if err!=nil {
		return false, err
	}

	return true, nil
}