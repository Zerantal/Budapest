package interfaces

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/model"
)

// UserRepository defines the methods for interacting with the User data.
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, userID primitive.ObjectID) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, upsert bool) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}
