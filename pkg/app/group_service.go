package app

import (
	"budapest/pkg/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GroupService provides business logic for managing groups.
type GroupService struct {
	groupRepository domain.GroupRepository
}

// NewGroupService creates a new GroupService with the given GroupRepository.
func NewGroupService(groupRepository domain.GroupRepository) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
	}
}

// CreateGroup creates a new group.
func (s *GroupService) CreateGroup(ctx context.Context, group *domain.Group) error {
	// Check if the group already exists by ID.
	existingGroup, err := s.groupRepository.GetGroupByID(ctx, group.ID)
	if err == nil && existingGroup != nil {
		return errors.New("group with the same ID already exists")
	}

	return s.groupRepository.CreateGroup(ctx, group)
}

// GetGroupByID retrieves a group by its ID.
func (s *GroupService) GetGroupByID(ctx context.Context, groupID primitive.ObjectID) (*domain.Group, error) {
	return s.groupRepository.GetGroupByID(ctx, groupID)
}

// UpdateGroup updates an existing group.
func (s *GroupService) UpdateGroup(ctx context.Context, group *domain.Group) error {
	return s.groupRepository.UpdateGroup(ctx, group)
}

// DeleteGroup deletes a group by its ID.
func (s *GroupService) DeleteGroup(ctx context.Context, groupID primitive.ObjectID) error {
	return s.groupRepository.DeleteGroup(ctx, groupID)
}
