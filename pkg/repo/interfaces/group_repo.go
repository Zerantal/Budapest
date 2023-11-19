package interfaces

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/model"
)

// GroupRepository defines the methods for interacting with the Group data.
type GroupRepository interface {
	CreateGroup(ctx context.Context, group *model.Group) error
	GetGroupByID(ctx context.Context, groupID primitive.ObjectID) (*model.Group, error)
	UpdateGroup(ctx context.Context, group *model.Group) error
	DeleteGroup(ctx context.Context, groupID primitive.ObjectID) error
}
