package renders

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//
var functions = template.FuncMap{}

// RenderTemplates
func RenderTemplates(w http.ResponseWriter, str string) {
	_, err := RenderTmplTest(w)
	if err != nil {
		log.Println("error getting template cache", err)
	}
	parseTmpl, _ := template.ParseFiles("./templates/" + str)
	err = parseTmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Test
func RenderTmplTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	// Define a template caches variable
	tmplCache := map[string]*template.Template{}
	//
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmplCache, err
	}

	//
	for _, page := range pages {
		pname := filepath.Base(page)
		fmt.Println("Page info", pname)
		ts, err := template.New(pname).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		//
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tmplCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tmplCache, err
			}
		}
		tmplCache[pname] = ts
	}
	return tmplCache, nil
}
