package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"stratus-api/color"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

func InitMongoDB() error {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	if DATABASE_URL=="" {
		return errors.New("DATABASE_URL not found")
	}

	clientOptions := options.Client().ApplyURI(DATABASE_URL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err!=nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	MongoDBClient = client

	fmt.Printf("%sConnected to MongoDB\n%s", color.Green, color.Reset)
	return nil
}