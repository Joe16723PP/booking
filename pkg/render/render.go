package render

import (
	"bytes"
	"fmt"
	"github.com/joe16723/booking/pkg/config"
	"github.com/joe16723/booking/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplate set the config for the template
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(writer http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	// check if dev mode load all template every time
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from template cahce")
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)
	// execute template to buffer
	_ = t.Execute(buffer, templateData)

	_, err := buffer.WriteTo(writer)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}

}

// CreateTemplateCache create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err := templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			myCache[name] = templateSet
		} else {
			myCache[name] = templateSet
		}
	}
	return myCache, nil
}
