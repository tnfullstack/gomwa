package main

import (
	"html/template"
	"log"
	"net/http"
)

// renderTemplates
func renderTemplates(w http.ResponseWriter, str string) {
	tmpl, _ := template.ParseFiles("templates/" + str)

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
