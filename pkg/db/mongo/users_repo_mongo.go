package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"budapest/pkg/domain"
)

func (r *MongoDBRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"_id": id}
	err := r.userCollection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoDBRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"username": username}
	err := r.userCollection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoDBRepository) UpdateUser(ctx context.Context, user *domain.User, upsert bool) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": user,
	}

	opts := options.Update()
	if upsert {
		opts.SetUpsert(true)
	}

	_, err := r.userCollection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *MongoDBRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.userCollection.DeleteOne(ctx, filter)
	return err
}
