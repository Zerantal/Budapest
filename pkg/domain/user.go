package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user entity in the CMS.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"` // Securely hashed password
	Salt     string             `bson:"salt"`     // Salt for password hashing
	Roles    []string           `bson:"roles"`    // User roles or permissions
	Groups   []string           `bson:"groups"`   // Groups that the user is a member of
}

// UserService defines methods to manage user entities.
type UserService interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *User) error

	// GetUserByID retrieves a user by ID.
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)

	// GetUserByUsername retrieves a user by username.
	GetUserByUsername(ctx context.Context, username string) (*User, error)

	// UpdateUser updates an existing user.
	UpdateUser(ctx context.Context, user *User, upsert bool) error

	// DeleteUser deletes a user by ID.
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}

// UserRepository defines database methods for user entities.
type UserRepository interface {
	// CreateUser creates a new user in the database.
	CreateUser(ctx context.Context, user *User) error

	// GetUserByID retrieves a user by ID from the database.
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*User, error)

	// GetUserByUsername retrieves a user by username from the database.
	GetUserByUsername(ctx context.Context, username string) (*User, error)

	// UpdateUser updates an existing user in the database.
	UpdateUser(ctx context.Context, user *User, upsert bool) error

	// DeleteUser deletes a user by ID from the database.
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}
