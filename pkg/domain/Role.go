package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role represents a user role or permission level in the CMS.
type Role struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"` // Name of the role (e.g., "Admin", "Editor", "Viewer")
}

// RoleRepository defines the methods for interacting with the Role data.
type RoleRepository interface {
	CreateRole(ctx context.Context, role *Role) error
	GetRoleByID(ctx context.Context, roleID primitive.ObjectID) (*Role, error)
	UpdateRole(ctx context.Context, role *Role) error
	DeleteRole(ctx context.Context, roleID primitive.ObjectID) error
}

// RoleService defines the business logic for managing Role entities.
type RoleService interface {
	CreateRole(ctx context.Context, role *Role) error
	GetRoleByID(ctx context.Context, roleID primitive.ObjectID) (*Role, error)
	UpdateRole(ctx context.Context, role *Role) error
	DeleteRole(ctx context.Context, roleID primitive.ObjectID) error
}
