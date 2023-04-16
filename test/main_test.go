package test

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/benweissmann/memongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InMemoryMongoDB struct {
	once        sync.Once
	mongoClient *mongo.Client
}

var instance *InMemoryMongoDB

// GetInMemoryMongoDBInstance returns a singleton instance of an in-memory MongoDB.
func GetInMemoryMongoDBInstance() *InMemoryMongoDB {
	if instance == nil {
		instance = &InMemoryMongoDB{}
	}
	return instance
}

// Connect sets up the in-memory MongoDB instance and connects the MongoDB client to it.
func (db *InMemoryMongoDB) Connect() {
	db.once.Do(func() {
		mongoServer, err := memongo.Start("latest")
		if err != nil {
			log.Fatalf("Failed to start in-memory MongoDB server: %v", err)
		}

		clientOptions := options.Client().ApplyURI(mongoServer.URI())
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatalf("Failed to create MongoDB client: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatalf("Failed to connect to in-memory MongoDB: %v", err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to ping in-memory MongoDB: %v", err)
		}

		db.mongoClient = client
	})
}

// Client returns the MongoDB client instance.
func (db *InMemoryMongoDB) Client() *mongo.Client {
	return db.mongoClient
}

// Disconnect disconnects the MongoDB client and stops the in-memory MongoDB server.
func (db *InMemoryMongoDB) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.mongoClient.Disconnect(ctx)
	if err != nil {
		log.Printf("Failed to disconnect from in-memory MongoDB: %v", err)
	}
}
