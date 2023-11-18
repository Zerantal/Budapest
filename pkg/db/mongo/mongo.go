package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB connects to a MongoDB database.
func ConnectToMongoDB(ctx context.Context, URI string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// CloseMongoDBConnection closes the MongoDB client.
func CloseMongoDBConnection(client *mongo.Client) {
	if client != nil {
		client.Disconnect(context.Background())
	}
}

// MongoDBRepository represents a MongoDB implementation of the UserRepository interface.
type MongoDBRepository struct {
	client          *mongo.Client
	userCollection  *mongo.Collection
	roleCollection  *mongo.Collection
	groupCollection *mongo.Collection
}

func NewMongoDBRepository(client *mongo.Client, dbName string) *MongoDBRepository {
	db := client.Database(dbName)
	userCollection := db.Collection("users")
	roleCollection := db.Collection("roles")
	groupCollection := db.Collection("groups")
	return &MongoDBRepository{
		client:          client,
		userCollection:  userCollection,
		roleCollection:  roleCollection,
		groupCollection: groupCollection,
	}
}
