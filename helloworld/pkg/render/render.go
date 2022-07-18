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
	// create a template cache
	tmap, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get request template from cache
	t, ok := tmap[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// parseTmpl, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = parseTmpl.Execute(w, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func createTemplateCache() (map[string]*template.Template, error) {
	// tmplCache := make(map[string]*template.Template)
	tmplCache := map[string]*template.Template{}

	// get all of the files name *.html from ./templates
	pages, err := filepath.Glob("./templates/*.html")
	fmt.Println(pages)
	if err != nil {
		return tmplCache, err
	}

	// range through all file-paths ending with *.html
	for _, page := range pages {
		name := filepath.Base(page)

		// get the pointer to the template pages
		ts, err := template.New(name).ParseFiles(page)
		fmt.Println(ts)
		if err != nil {
			return tmplCache, err
		}

		// get the template file name
		matches, err := filepath.Glob("./templates/*.html")
		fmt.Println(matches)
		if err != nil {
			return tmplCache, err
		}

		if len(matches) > 0 {
			// get the pointer to template
			ts, err = ts.ParseGlob("./templates/*.html")
			if err != nil {
				return tmplCache, err
			}
		}
		// mapping the templates filename to the pointers to the template ParseFiles
		tmplCache[name] = ts
	}
	return tmplCache, nil
}
