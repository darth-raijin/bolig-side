package test

import (
	"context"
	"fmt"

	"github.com/benweissmann/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func ConnectToMockMongoDB() (*mongo.Client, error) {
	// Create a new mock MongoDB server
	server, err := memongo.Start("4.4")
	if err != nil {
		return nil, fmt.Errorf("failed to start mock MongoDB server: %v", err)
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(server.URI())

	// Connect to the mock MongoDB server
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mock MongoDB server: %v", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mock MongoDB server: %v", err)
	}

	fmt.Println("Connected to mock MongoDB server!")

	return client, nil
}
