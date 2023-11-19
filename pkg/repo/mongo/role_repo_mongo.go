package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"budapest/pkg/model"
)

func (r *MongoDBRepository) CreateRole(ctx context.Context, role *model.Role) error {
	_, err := r.roleCollection.InsertOne(ctx, role)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) GetRoleByID(ctx context.Context, roleID primitive.ObjectID) (*model.Role, error) {
	var role model.Role
	err := r.roleCollection.FindOne(ctx, bson.M{"_id": roleID}).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *MongoDBRepository) GetRoleByName(ctx context.Context, name string) (*model.Role, error) {
	var role model.Role
	err := r.roleCollection.FindOne(ctx, bson.M{"name": name}).Decode(&role)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *MongoDBRepository) UpdateRole(ctx context.Context, role *model.Role) error {
	update := bson.M{"$set": role}
	_, err := r.roleCollection.UpdateOne(ctx, bson.M{"_id": role.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) DeleteRole(ctx context.Context, roleID primitive.ObjectID) error {
	_, err := r.roleCollection.DeleteOne(ctx, bson.M{"_id": roleID})
	if err != nil {
		return err
	}
	return nil
}
