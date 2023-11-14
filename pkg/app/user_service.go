package app

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/domain"
)

// UserService represents a service for managing user entities.
type UserService struct {
	userRepo domain.UserRepository
}

// NewUserService creates a new instance of UserService with the given UserRepository.
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

// CreateUser creates a new user in the system.
func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	// Validate user data if needed
	if user == nil {
		return errors.New("user data is nil")
	}

	// Add any additional validation logic here

	// Create the user in the repository
	return s.userRepo.CreateUser(ctx, user)
}

// GetUserByID retrieves a user by their unique ID.
func (s *UserService) GetUserByID(ctx context.Context, userID primitive.ObjectID) (*domain.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

// GetUserByID retrieves a user by their unique ID.
func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.userRepo.GetUserByUsername(ctx, username)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(ctx context.Context, updatedUser *domain.User, upsert bool) error {
	// Validate user data if needed
	if updatedUser == nil {
		return errors.New("updated user data is nil")
	}

	// Add any additional validation logic here

	// Update the user in the repository
	return s.userRepo.UpdateUser(ctx, updatedUser, upsert)
}

// DeleteUser deletes a user by their unique ID.
func (s *UserService) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	return s.userRepo.DeleteUser(ctx, userID)
}
