package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/config"
	"budapest/pkg/model"
	mongo_helper "budapest/pkg/repo/mongo"
	"budapest/pkg/service"
)

func seedAdminUser(appService *service.AppService) error {
	// Create a context for the database operation
	ctx := context.TODO() // Replace with an appropriate context if needed

	// Define the admin user data
	adminUser := &model.User{
		ID:       primitive.NewObjectID(),
		Username: "admin",
		Email:    "admin@example.com",
		Password: "adminpassword", // Password should be securely hashed in a real scenario
	}

	// Create the admin user using the UserService
	if err := appService.UserService.CreateUser(ctx, adminUser); err != nil {
		return err
	}

	fmt.Println("Admin user seeded successfully.")
	return nil
}

func main() {
	// Create a top-level context for the application
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Define a command-line flag for the configuration file path
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config.json", "Path to the configuration file")
	flag.Parse()

	// If the provided configFilePath is not an absolute path, make it relative to the current directory
	if !filepath.IsAbs(configFilePath) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get the current working directory: %v", err)
		}
		configFilePath = filepath.Join(wd, configFilePath)
	}

	// Load the configuration from the specified JSON file
	cfg, err := config.LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the MongoDB client
	client, err := mongo_helper.ConnectToMongoDB(context.Background(), cfg.DatabaseConfig.ConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongo_helper.CloseMongoDBConnection(client)

	// Create an instance of AppService
	appService, err := service.NewAppService(ctx, cfg.DatabaseConfig)
	if err != nil {
		log.Fatalf("Failed to initialize AppService: %v", err)
	}

	// Seed the admin user
	if err := seedAdminUser(appService); err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	fmt.Println("Database seeding complete.")
}
