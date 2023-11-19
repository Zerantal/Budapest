package seeder

import (
	"context"
	"log"

	"budapest/pkg/model"
	"budapest/pkg/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func seedAdminUser(ctx context.Context, appService *service.AppService) error {

	// Check if admin user already exists
	existingAdminUser, err := appService.UserService.GetUserByUsername(ctx, "admin")
	if err != nil {
		return err
	}

	adminUser := &model.User{
		ID:       primitive.NewObjectID(),
		Username: "admin",
		Email:    "admin@example.com",
		Password: "adminpasssword", // Remember to hash the password
	}

	if existingAdminUser != nil {
		adminUser.ID = existingAdminUser.ID
		if err := appService.UserService.UpdateUser(ctx, adminUser); err != nil {
			return err
		}
		log.Println("Admin user updated successfully.")
	} else {
		if err := appService.UserService.CreateUser(ctx, adminUser); err != nil {
			return err
		}
		log.Println("Admin user created successfully.")
	}

	return nil
}
