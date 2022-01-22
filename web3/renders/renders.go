package renders

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tvn9/gomwa/web3/config"
	"github.com/tvn9/gomwa/web3/models"
)

// functions
// var functions = template.FuncMap{}

// Set variable for AppConfig
var appCfg *config.AppConfig

// NewTemplate
func NewTemplates(a *config.AppConfig) {
	appCfg = a
}

// renderTemplates
func RenderTemplates(w http.ResponseWriter, str string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// Get templates cache from AppConfig struct
	if appCfg.UseCache {
		tc = appCfg.TmplCache
	} else {
		tc, _ = CreateTmplCache()
	}

	t, ok := tc[str]
	if !ok {
		log.Fatal("could not get template from tmplCache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTmplCache
func CreateTmplCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	// pages, err := filepath.Glob("./templates/*.page.html")
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return tmplCache, err
	}

	for _, p := range pages {
		name := filepath.Base(p)
		// ts, err := template.New(name).Funcs(functions).ParseFiles(p)
		ts, err := template.New(name).ParseFiles(p)
		if err != nil {
			return tmplCache, err
		}

		// matches, err := filepath.Glob("./templates/*.layout.html")
		// if err != nil {
		// 	return tmplCache, err
		// }
		// // fmt.Println(matches)

		//if len(pages) > 0 {
		ts, err = ts.ParseGlob("./templates/*.layout.html")
		if err != nil {
			return tmplCache, err
		}

		tmplCache[name] = ts
	}
	fmt.Println("From CreateTmplCache func:", tmplCache)
	return tmplCache, nil
}
