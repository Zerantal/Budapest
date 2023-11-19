package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"budapest/pkg/model"
	repo_interface "budapest/pkg/repo/interfaces"
	"budapest/pkg/service/interfaces"
)

// webPageTemplateService implements the WebPageTemplateService interface
type webPageTemplateService struct {
	repo repo_interface.WebPageTemplateRepository
}

// NewWebPageTemplateService creates a new instance of WebPageTemplateService
func NewWebPageTemplateService(repo repo_interface.WebPageTemplateRepository) interfaces.WebPageTemplateService {
	return &webPageTemplateService{repo: repo}
}

// CreateTemplate creates a new web page template
func (s *webPageTemplateService) CreateTemplate(ctx context.Context, template *model.WebPageTemplate) error {
	return s.repo.CreateTemplate(ctx, template)
}

// GetTemplateByID retrieves a template by its ID
func (s *webPageTemplateService) GetTemplateByID(ctx context.Context, id primitive.ObjectID) (*model.WebPageTemplate, error) {
	return s.repo.GetTemplateByID(ctx, id)
}

// GetTemplateByName retrieves a template by its name
func (s *webPageTemplateService) GetTemplateByName(ctx context.Context, name string) (*model.WebPageTemplate, error) {
	return s.repo.GetTemplateByName(ctx, name)
}

// UpdateTemplate updates an existing template
func (s *webPageTemplateService) UpdateTemplate(ctx context.Context, updatedTemplate *model.WebPageTemplate) error {
	return s.repo.UpdateTemplate(ctx, updatedTemplate)
}

// DeleteTemplate deletes a template by its ID
func (s *webPageTemplateService) DeleteTemplate(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.DeleteTemplate(ctx, id)
}
