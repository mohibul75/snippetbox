package main

import(
	"github.com/mohibul75/snippetbox-go-project/internal/models"
)

type templateData struct{
	Snippet *models.Snippet
	Snippets [] *models.Snippet
}