package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user entity in the CMS.
type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Username string               `bson:"username"`
	Email    string               `bson:"email"`
	Password string               `bson:"password"` // Securely hashed password
	Salt     string               `bson:"salt"`     // Salt for password hashing
	Roles    []primitive.ObjectID `bson:"roles"`    // User roles or permissions
	Groups   []primitive.ObjectID `bson:"groups"`   // Groups that the user is a member of
}
