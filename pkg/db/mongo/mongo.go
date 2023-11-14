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
	client         *mongo.Client
	userCollection *mongo.Collection
}

func NewMongoDBRepository(client *mongo.Client, dbName, collectionName string) *MongoDBRepository {
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &MongoDBRepository{
		client:         client,
		userCollection: collection,
	}
}
