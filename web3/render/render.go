package render

import (
	"html/template"
	"log"
	"net/http"
)

// renderTemplates
func RenderTemplates(w http.ResponseWriter, str string) {
	tmpl, _ := template.ParseFiles("templates/" + str)

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
