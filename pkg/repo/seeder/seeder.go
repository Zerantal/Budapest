package seeder

import (
	"context"

	"budapest/pkg/service"
)

func SeedEntities(ctx context.Context, appService *service.AppService) error {
	if err := seedAdminUser(ctx, appService); err != nil {
		return err
	}

	if err := SeedWebPageTemplates(ctx, appService); err != nil {
		return err
	}

	return nil
}
