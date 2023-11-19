package seeder

import (
	"context"

	"budapest/pkg/service"
)

func SeedEntities(ctx context.Context, appService *service.AppService) error {
	if err := seedAdminUser(ctx, appService); err != nil {
		return err
	}

	// Additional seeds for other entities can be added here
	return nil
}
