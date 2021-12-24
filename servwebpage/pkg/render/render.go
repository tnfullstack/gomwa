package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Writing template browser:", err)
	}

	// parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	// err = parseTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error parsing template:", err)
	// 	return
	// }
}

// CreateTemplateCache creates a map of template
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		log.Println(err)
		return myCache, err
	}

	for _, p := range pages {
		name := filepath.Base(p)

		// fmt.Println("Page is currently", p)
		ts, err := template.New(name).Funcs(functions).ParseFiles(p)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
