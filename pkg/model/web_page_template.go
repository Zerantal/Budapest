package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// WebPageTemplate represents a template for web pages.
type WebPageTemplate struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Content string             `bson:"html_content"`
}
