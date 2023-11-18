package domain

import (
	"context"

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

// UserRepository defines the methods for interacting with the User data.
type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, userID primitive.ObjectID) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, user *User, upsert bool) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}

// UserService defines the business logic for managing User entities.
type UserService interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, userID primitive.ObjectID) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, user *User, upsert bool) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}
