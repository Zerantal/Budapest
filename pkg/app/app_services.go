package app

import (
	"context"

	"budapest/pkg/db"
	mongo_helper "budapest/pkg/db/mongo"
	"budapest/pkg/domain"
)

type AppService struct {
	UserService domain.UserService
}

// NewAppService creates a new instance of AppService with the specified database configuration.
// returns services that uses mongo (could be extended to use other db engines)
func NewAppService(ctx context.Context, dbConfig db.DatabaseConfig) (*AppService, error) {
	// Initialize MongoDB client using the provided configuration
	client, err := mongo_helper.ConnectToMongoDB(ctx, dbConfig.ConnectionString)
	if err != nil {
		return nil, err
	}

	// Create repositories and services
	userRepo := mongo_helper.NewMongoDBRepository(client, dbConfig.DatabaseName, "Users")

	userService := NewUserService(userRepo)

	// Create the AppService instance
	appService := &AppService{
		UserService: userService}

	return appService, nil
}
