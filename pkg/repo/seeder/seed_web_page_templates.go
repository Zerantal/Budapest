package seeder

import (
	"context"
	"log"

	"budapest/pkg/model"
	"budapest/pkg/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedWebPageTemplates(ctx context.Context, appService *service.AppService) error {

	// List of templates to seed
	templates := []model.WebPageTemplate{
		{
			Name:    "Home Page Template",
			Content: "<html><body>Home Page Content</body></html>",
		},
		{
			Name:    "Contact Page Template",
			Content: "<html><body>Contact Page Content</body></html>",
		},
	}

	for _, tmpl := range templates {
		err := seedTemplate(ctx, appService, &tmpl)
		if err != nil {
			log.Printf("Failed to seed template '%s': %v\n", tmpl.Name, err)
			return err
		}
	}

	return nil
}

func seedTemplate(ctx context.Context, appService *service.AppService, template *model.WebPageTemplate) error {
	// Try to retrieve the template by its name (assuming you have a method for this)
	// For the sake of this example, I'm assuming a method like 'GetTemplateByName' exists
	existingTemplate, err := appService.WebPageTemplateService.GetTemplateByName(ctx, template.Name)
	if err != nil {
		return err
	}

	if existingTemplate != nil {
		// Template already exists, update it
		template.ID = existingTemplate.ID
		if err := appService.WebPageTemplateService.UpdateTemplate(ctx, template); err != nil {
			return err
		}
		log.Printf("Template '%s' updated successfully.\n", template.Name)
	} else {
		// Template does not exist, create it
		template.ID = primitive.NewObjectID() // Assign a new ID
		if err := appService.WebPageTemplateService.CreateTemplate(ctx, template); err != nil {
			return err
		}
		log.Printf("Template '%s' created successfully.\n", template.Name)
	}
	return nil
}
