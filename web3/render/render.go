package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// functions
var functions = template.FuncMap{}

// renderTemplates
func RenderTemplates(w http.ResponseWriter, str string) {
	tc, err := CreateTmplCache()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("from RenderTemplates func:", tc)

	t, ok := tc[str]
	if !ok {
		log.Fatal(err)
	}
	// fmt.Println("look in tc map for maching key:", str, t)

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTmplCache
func CreateTmplCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return tmplCache, err
	}
	// fmt.Println(pages)

	for _, p := range pages {
		name := filepath.Base(p)
		// fmt.Println(p)
		ts, err := template.New(name).Funcs(functions).ParseFiles(p)
		if err != nil {
			return tmplCache, err
		}
		//fmt.Println(ts)

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return tmplCache, err
		}
		// fmt.Println(matches)

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			// fmt.Println(ts)
			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[name] = ts
	}
	// fmt.Println("From CreateTmplCache func:", tmplCache)
	return tmplCache, nil
}
