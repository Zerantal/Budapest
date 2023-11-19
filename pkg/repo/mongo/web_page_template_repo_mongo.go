package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"budapest/pkg/model"
)

// CreateTemplate creates a new web page template
func (r *MongoDBRepository) CreateTemplate(ctx context.Context, template *model.WebPageTemplate) error {
	_, err := r.WebPageTemplateCollection.InsertOne(ctx, template)
	return err
}

// GetTemplateByID retrieves a template by its ID
func (r *MongoDBRepository) GetTemplateByID(ctx context.Context, id primitive.ObjectID) (*model.WebPageTemplate, error) {
	var template model.WebPageTemplate
	filter := bson.M{"_id": id}
	err := r.WebPageTemplateCollection.FindOne(ctx, filter).Decode(&template)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("template not found")
	}
	return &template, err
}

// GetTemplateByName retrieves a template by its name
func (r *MongoDBRepository) GetTemplateByName(ctx context.Context, name string) (*model.WebPageTemplate, error) {
	var template model.WebPageTemplate
	filter := bson.M{"name": name}
	err := r.WebPageTemplateCollection.FindOne(ctx, filter).Decode(&template)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("template not found")
	}
	return &template, err
}

// UpdateTemplate updates an existing template
func (r *MongoDBRepository) UpdateTemplate(ctx context.Context, updatedTemplate *model.WebPageTemplate) error {
	filter := bson.M{"_id": updatedTemplate.ID}
	update := bson.M{"$set": updatedTemplate}
	_, err := r.WebPageTemplateCollection.UpdateOne(ctx, filter, update)
	return err
}

// DeleteTemplate deletes a template by its ID
func (r *MongoDBRepository) DeleteTemplate(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.WebPageTemplateCollection.DeleteOne(ctx, filter)
	return err
}
