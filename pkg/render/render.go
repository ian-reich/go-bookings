package render

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	// Create template cache
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get Request Template From Cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	// Render the template

	parsedTemplate, err := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		fmt.Println("error parsing template:", err)
		return
	}

	err = parsedTemplate.Execute(w, td)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Println("error executing template:", err)
		return
	}
}

// V1.0
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the named files *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through al files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
