// WebPageTemplateRepository defines the database methods for web page templates.
package interfaces

import (
	"context"

	"budapest/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebPageTemplateRepository interface {
	CreateTemplate(ctx context.Context, template *model.WebPageTemplate) error
	GetTemplateByID(ctx context.Context, id primitive.ObjectID) (*model.WebPageTemplate, error)
	GetTemplateByName(ctx context.Context, name string) (*model.WebPageTemplate, error)
	UpdateTemplate(ctx context.Context, updatedTemplate *model.WebPageTemplate) error
	DeleteTemplate(ctx context.Context, id primitive.ObjectID) error
}
