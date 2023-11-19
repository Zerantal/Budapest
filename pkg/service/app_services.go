package service

import (
	"context"

	"budapest/pkg/repo"
	mongo_repo "budapest/pkg/repo/mongo"
	"budapest/pkg/service/interfaces"
)

type AppService struct {
	UserService  interfaces.UserService
	RoleService  interfaces.RoleService
	GroupService interfaces.GroupService
}

// NewAppService creates a new instance of AppService with the specified database configuration.
// returns services that uses mongo (could be extended to use other db engines)
func NewAppService(ctx context.Context, dbConfig repo.DatabaseConfig) (*AppService, error) {
	// Initialize MongoDB client using the provided configuration
	client, err := mongo_repo.ConnectToMongoDB(ctx, dbConfig.ConnectionString)
	if err != nil {
		return nil, err
	}

	// Create repositories and services
	repo := mongo_repo.NewMongoDBRepository(client, dbConfig.DatabaseName)

	userService := NewUserService(repo)
	RoleService := NewRoleService(repo)
	GroupService := NewGroupService(repo)

	// Create the AppService instance
	appService := &AppService{
		UserService:  userService,
		RoleService:  RoleService,
		GroupService: GroupService}

	return appService, nil
}
