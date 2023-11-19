package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role represents a user role or permission level in the CMS.
type Role struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"` // Name of the role (e.g., "Admin", "Editor", "Viewer")
}
