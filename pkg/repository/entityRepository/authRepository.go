package entityrepository

import (
	"context"
	"errors"
	"time"

	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"
	"github.com/darth-raijin/bolig-side/api/models/entities"
	"github.com/darth-raijin/bolig-side/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUser(user entities.User) (entities.User, error) {
	// Connect to MongoDB
	client, collection, err := repository.MongoConnectDatabase("Users")
	if err != nil {
		return entities.User{}, err
	}

	defer repository.DisconnectMongo(context.Background(), client)

	// Validate email is unique
	filter := bson.M{"email": user.Email}
	var existingUser entities.User
	err = collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == nil {
		return entities.User{}, errors.New(errorDto.EmailNotUnique.Message)
	}

	// Generate ID and set timestamps
	user.ID = primitive.NewObjectID()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Insert user to the database
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string, password string) (entities.User, error) {
	// Connect to MongoDB
	client, collection, err := repository.MongoConnectDatabase("Users")
	if err != nil {
		return entities.User{}, err
	}

	defer repository.DisconnectMongo(context.Background(), client)

	// Find user by email and hashed password
	filter := bson.M{"email": email, "password": password}
	var foundUser entities.User
	err = collection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.User{}, errors.New(errorDto.UserNotFound.Message)
		}
		return entities.User{}, err
	}

	return foundUser, nil
}
