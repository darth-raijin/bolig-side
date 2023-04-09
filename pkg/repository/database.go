package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/darth-raijin/bolig-side/pkg/utility"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func MongoConnectDatabase() {
	var err error

	fmt.Print("foo")
	fmt.Print(utility.GetAppConfig().Database)
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%v:%v@%v.mongodb.net/?retryWrites=true&w=majority",
		utility.GetAppConfig().Database.Username,
		utility.GetAppConfig().Database.Password,
		utility.GetAppConfig().Database.Database,
	))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Disconnect from MongoDB
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB!")
}
