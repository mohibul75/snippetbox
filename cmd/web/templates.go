package main

import (
	"html/template"
	"path/filepath"

	"github.com/mohibul75/snippetbox-go-project/internal/models"
)

type templateData struct{
	Snippet *models.Snippet
	Snippets [] *models.Snippet
}

func newTemplateCache()(map[string]*template.Template, error){

	cache:= map[string]*template.Template{}

	pages, err:= filepath.Glob("./ui/html/pages/*.tmpl")

	if err!=nil{
		return nil, err
	}

	for _, page := range pages {

		name:= filepath.Base(page)


		files:= []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		ts,err:= template.ParseFiles(files...)
		if err!=nil {
			return nil, err
		}

		// ts,err:=template.ParseFiles("./ui/html/base.tmpl")
		// if err!=nil{
		// 	return nil,err
		// }

		// ts, err= template.ParseGlob("./ui/html/partials/*.tmpl")
		// if err!=nil{
		// 	return nil,err
		// }
	
		// ts, err= template.ParseFiles(page)
		// if err!=nil{
		// 	return nil,err
		// }
		cache[name]=ts

  }
	return cache, nil
}