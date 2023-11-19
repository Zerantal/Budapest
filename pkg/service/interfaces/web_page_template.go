package interfaces

import (
	"context"

	"budapest/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WebPageTemplateService defines the methods to manage web page templates.
type WebPageTemplateService interface {
	CreateTemplate(ctx context.Context, template *model.WebPageTemplate) error
	GetTemplateByID(ctx context.Context, id primitive.ObjectID) (*model.WebPageTemplate, error)
	GetTemplateByName(ctx context.Context, name string) (*model.WebPageTemplate, error)
	UpdateTemplate(ctx context.Context, updatedTemplate *model.WebPageTemplate) error
	DeleteTemplate(ctx context.Context, id primitive.ObjectID) error
}
