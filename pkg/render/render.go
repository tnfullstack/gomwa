package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tvn9/gomwa/pkg/config"
	"github.com/tvn9/gomwa/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.AppData) *models.AppData {

	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.AppData) {

	var tc map[string]*template.Template
	if app.UseCache {
		log.Println("Using templates cache.")
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatalf("fails to create template cache.")
		}
		log.Println("Recreate template evertime accessing a web page.")
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalf("fails to access template cache.")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	if err := t.Execute(buf, td); err != nil {
		log.Println(err)
	}

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tCache := map[string]*template.Template{}

	// get all o the templates file name from ./templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return tCache, err
	}

	// range through all files ending with *.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return tCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return tCache, err
			}
		}

		tCache[name] = ts
	}
	return tCache, nil
}

// Below is code written using basic cache concept, the above is the improve version.

/*
var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	templates, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	err := templates.Execute(w, nil)
	if err != nil {
		log.Fatalf("fails to execute template %v\n", err)
	}
}

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check template exist in map
	_, inMap := tc[t]
	if !inMap {
		// create template map
		log.Println("Creating template and adding to cache")
		err := createTemplateCache(t)
		if err != nil {
			log.Fatalf("fails to create template cache. %v\n", err)
		}
	} else {
		// using template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("fails to execute template. %v\n", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		"./templates/" + t,
		"./templates/base.html",
	}

	// parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
*/
