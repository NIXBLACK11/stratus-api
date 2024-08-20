package actions

import (
	"context"
	"errors"
	"os"
	"stratus-api/database"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProjects(username string) ([]string, error) {
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	if DATABASE_NAME == "" {
		return nil, errors.New("Database name not present")
	}

	if database.MongoDBClient == nil {
		return nil, errors.New("MongoDB client is not initialized")
	}

	projectsCollection := database.MongoDBClient.Database(DATABASE_NAME).Collection("Projects")

	cursor, err := projectsCollection.Find(context.TODO(), bson.D{
        {Key: "username", Value: username},
    })
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var projects []string
    for cursor.Next(context.TODO()) {
        var project bson.M
        if err := cursor.Decode(&project); err != nil {
            return nil, err
        }
        // Extract project name (assuming "projectname" field in the document)
        if projectName, ok := project["projectname"].(string); ok {
            projects = append(projects, projectName)
        }
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return projects, nil

}