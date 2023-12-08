package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get the requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	tCache := map[string]*template.Template{} // other way declare map

	// get all of the templates file from the ./templates directory
	tmplPages, err := filepath.Glob("./templates/*.tmpl.html")
	fmt.Println(tmplPages)
	if err != nil {
		return tCache, err
	}

	// range through all files ending with *.tmpl.html
	for _, tp := range tmplPages {
		tmplName := filepath.Base(tp)
		ts, err := template.New(tmplName).ParseFiles(tp)
		if err != nil {
			return tCache, err
		}

		tmplSlice, err := filepath.Glob("./templates/*.tmpl.html")
		if err != nil {
			log.Fatal(err)
		}

		if len(tmplSlice) > 0 {
			ts, err = ts.ParseGlob("./templates/*.tmpl.html")
			if err != nil {
				return tCache, err
			}
		}
		tCache[tmplName] = ts
	}
	return tCache, nil
}

////// The below bock of code is one method of create template cache //////

/*
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in the cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("Creating template and adding to cache")
		// need to create the template
		err = createTemplateCache(t)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// we have the template in the cache
		log.Println("Using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/homelayout.html",
	}

	// parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Fatal(err)
	}

	// store the parsed templete in map
	tc[t] = tmpl

	return nil
}
*/
