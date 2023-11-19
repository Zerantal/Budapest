package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"budapest/pkg/model"
)

// CreateGroup creates a new group in the MongoDB database.
func (r *MongoDBRepository) CreateGroup(ctx context.Context, group *model.Group) error {
	_, err := r.groupCollection.InsertOne(ctx, group)
	if err != nil {
		return err
	}
	return nil
}

// GetGroupByID retrieves a group by its ObjectID.
func (r *MongoDBRepository) GetGroupByID(ctx context.Context, groupID primitive.ObjectID) (*model.Group, error) {
	var group model.Group
	err := r.groupCollection.FindOne(ctx, bson.M{"_id": groupID}).Decode(&group)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("group not found")
		}
		return nil, err
	}
	return &group, nil
}

// GetGroupByName retrieves a group by its name.
func (r *MongoDBRepository) GetGroupByName(ctx context.Context, name string) (*model.Group, error) {
	var group model.Group
	err := r.groupCollection.FindOne(ctx, bson.M{"name": name}).Decode(&group)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("group not found")
		}
		return nil, err
	}
	return &group, nil
}

// UpdateGroup updates an existing group in the MongoDB database.
func (r *MongoDBRepository) UpdateGroup(ctx context.Context, group *model.Group) error {
	update := bson.M{"$set": group}
	_, err := r.groupCollection.UpdateOne(ctx, bson.M{"_id": group.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGroup deletes a group by its ObjectID.
func (r *MongoDBRepository) DeleteGroup(ctx context.Context, groupID primitive.ObjectID) error {
	_, err := r.groupCollection.DeleteOne(ctx, bson.M{"_id": groupID})
	if err != nil {
		return err
	}
	return nil
}
