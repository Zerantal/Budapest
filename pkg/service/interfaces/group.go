package interfaces

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/model"
)

// GroupService defines the business logic for managing Group entities.
type GroupService interface {
	CreateGroup(ctx context.Context, group *model.Group) error
	GetGroupByID(ctx context.Context, groupID primitive.ObjectID) (*model.Group, error)
	UpdateGroup(ctx context.Context, group *model.Group) error
	DeleteGroup(ctx context.Context, groupID primitive.ObjectID) error
}
