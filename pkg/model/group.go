package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Group represents a group entity to enumerate users belonging to that group.
type Group struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Members []string           `bson:"members"`
}
