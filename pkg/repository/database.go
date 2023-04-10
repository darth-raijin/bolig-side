package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/darth-raijin/bolig-side/pkg/utility"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnectDatabase(collectionName string) (*mongo.Client, *mongo.Collection, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%v:%v@%v.mongodb.net/?retryWrites=true&w=majority",
		utility.GetAppConfig().Database.Username,
		utility.GetAppConfig().Database.Password,
		utility.GetAppConfig().Database.Database,
	))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Connected to MongoDB!")

	// Get the collection
	db := client.Database(utility.GetAppConfig().Database.Database)
	collection := db.Collection(collectionName)

	return client, collection, nil
}

func DisconnectMongo(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB!")
}
