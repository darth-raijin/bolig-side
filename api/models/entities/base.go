package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}
