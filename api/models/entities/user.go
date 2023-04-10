package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseEntity
	FirstName      string             `bson:"firstname,omitempty"`
	LastName       string             `bson:"lastname,omitempty"`
	Email          string             `bson:"email,omitempty"`
	Country        string             `bson:"country,omitempty"`
	Password       string             `bson:"password,omitempty"`
	Realtor        bool               `bson:"realtor,omitempty"`
	SubscriptionID primitive.ObjectID `bson:"subscriptionId,omitempty"`
	LastLoggedIn   time.Time          `bson:"lastLoggedIn,omitempty"`
}
