package main

import (
	"html/template"
	"megatroncodrr/snippetbox/pkg/models"
	"path/filepath"
)

// struct to hold structure for dynamic data that we want to pass to our html templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func cacheTemplate(dir string) (map[string]*template.Template, error) {
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	cache := map[string]*template.Template{}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil

}
