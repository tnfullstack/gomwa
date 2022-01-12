package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// var temp *template.Template

// func init() {
// 	temp = template.Must(template.ParseGlob("templates/*.tmpl"))
// }

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	// ListionAndServe
	_ = http.ListenAndServe(":9999", nil)

}

// HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplates(w, "index.page.tmpl")
	// err := temp.ExecuteTemplate(w, "index.page.tmpl", nil)
	// if err != nil {
	// 	log.Printf("Fails to execute template %v\n", err)
	// }
}

// AboutHandler
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
	// err := temp.ExecuteTemplate(w, "about.page.tmpl", nil)
	// if err != nil {
	// 	log.Printf("Fails to execute template %v\n", err)
	// }
}

func renderTemplates(w http.ResponseWriter, tname string) {
	tmplParse, _ := template.ParseFiles("templates/*page.tmpl" + tname)

	err := tmplParse.Execute(w, nil)
	if err != nil {
		fmt.Println("Can't execute template", err)
	}

}
