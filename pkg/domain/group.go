package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Group represents a group entity to enumerate users belonging to that group.
type Group struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Members []string           `bson:"members"`
}

// GroupRepository defines the methods for interacting with the Group data.
type GroupRepository interface {
	CreateGroup(ctx context.Context, group *Group) error
	GetGroupByID(ctx context.Context, groupID primitive.ObjectID) (*Group, error)
	UpdateGroup(ctx context.Context, group *Group) error
	DeleteGroup(ctx context.Context, groupID primitive.ObjectID) error
}

// GroupService defines the business logic for managing Group entities.
type GroupService interface {
	CreateGroup(ctx context.Context, group *Group) error
	GetGroupByID(ctx context.Context, groupID primitive.ObjectID) (*Group, error)
	UpdateGroup(ctx context.Context, group *Group) error
	DeleteGroup(ctx context.Context, groupID primitive.ObjectID) error
}
