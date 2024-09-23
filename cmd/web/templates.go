package main

import (
	"megatroncodrr/snippetbox/pkg/models"
)

// struct to hold structure for dynamic data that we want to pass to our html templates
type templateData struct {
	Snippet *models.Snippet
}
