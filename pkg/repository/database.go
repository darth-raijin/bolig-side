package repository

import (
	"context"
	"fmt"

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
		utility.Log(utility.ERROR, fmt.Sprintf("Failed to connect to database: %s", err))
		return nil, nil, err
	}

	// Get the collection
	db := client.Database(utility.GetAppConfig().Database.Database)
	collection := db.Collection(collectionName)
	utility.Log(utility.INFO, fmt.Sprintf("Successfully fetched collection: %s", collectionName))

	return client, collection, nil
}

func DisconnectMongo(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		utility.Log(utility.ERROR, fmt.Sprintf("Failed to disconnect to database: %s", err))
	}

	utility.Log(utility.ERROR, fmt.Sprintf("Failed to connect to database: %s", err))
}
