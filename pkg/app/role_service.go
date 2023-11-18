package app

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/domain"
)

// RoleService provides business logic for managing roles.
type RoleService struct {
	roleRepository domain.RoleRepository
}

// NewRoleService creates a new RoleService with the given RoleRepository.
func NewRoleService(roleRepository domain.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

// CreateRole creates a new role.
func (s *RoleService) CreateRole(ctx context.Context, role *domain.Role) error {
	return s.roleRepository.CreateRole(ctx, role)
}

// GetRoleByID retrieves a role by its ID.
func (s *RoleService) GetRoleByID(ctx context.Context, roleID primitive.ObjectID) (*domain.Role, error) {
	return s.roleRepository.GetRoleByID(ctx, roleID)
}

// UpdateRole updates an existing role.
func (s *RoleService) UpdateRole(ctx context.Context, role *domain.Role) error {
	return s.roleRepository.UpdateRole(ctx, role)
}

// DeleteRole deletes a role by its ID.
func (s *RoleService) DeleteRole(ctx context.Context, roleID primitive.ObjectID) error {
	return s.roleRepository.DeleteRole(ctx, roleID)
}
