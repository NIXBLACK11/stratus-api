package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"
	"stratus-api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProjectDetails(username string, projectname string) (models.Project, error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return models.Project{}, errors.New("Database name not present")
	}

	if database.MongoDBClient == nil {
		return models.Project{}, errors.New("MongoDB client is not initialized")
	}

	projectsCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Projects")

	var project models.Project
	err := projectsCollection.FindOne(context.TODO(), bson.D{
		{Key: "username", Value: username},
		{Key: "projectname", Value: projectname},
	}).Decode(&project)
	if err != nil {
		return models.Project{}, err
	}

	if project.UserName == "" && project.ProjectName == "" {
		return models.Project{}, errors.New("No such project")
	}

	return project, nil
}
