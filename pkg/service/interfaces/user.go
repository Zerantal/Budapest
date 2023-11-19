package interfaces

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/model"
)

// UserService defines the business logic for managing User entities.
type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, userID primitive.ObjectID) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, upsert bool) error
	DeleteUser(ctx context.Context, userID primitive.ObjectID) error
}
