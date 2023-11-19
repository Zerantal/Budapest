package interfaces

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/model"
)

// RoleService defines the business logic for managing Role entities.
type RoleService interface {
	CreateRole(ctx context.Context, role *model.Role) error
	GetRoleByID(ctx context.Context, roleID primitive.ObjectID) (*model.Role, error)
	UpdateRole(ctx context.Context, role *model.Role) error
	DeleteRole(ctx context.Context, roleID primitive.ObjectID) error
}
