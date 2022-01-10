package handlers

import (
    "log"
    "net/http"
    "text/template"
)

// declair tmpl object
var tmpl *template.Template

// init initialize the template variable
func init() {
    tmpl = template.Must(template.ParseGlob("../templates/*.tmpl"))
}

// Home returns the web app home page
func Home(w http.ResponseWriter, r *http.Request) {
    err := tmpl.ExecuteTemplate(w, "index.page.tmpl", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// About returns the web app about page
func About(w http.ResponseWriter, r *http.Request) {
    err := tmpl.ExecuteTemplate(w, "about.page.tmpl", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// Login returns the web app about page
func Login(w http.ResponseWriter, r *http.Request) {
    err := tmpl.ExecuteTemplate(w, "login.page.tmpl", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// Signup returns the web app about page
func Signup(w http.ResponseWriter, r *http.Request) {
    err := tmpl.ExecuteTemplate(w, "signup.page.tmpl", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// Contacts returns the web app about page
func Contacts(w http.ResponseWriter, r *http.Request) {
    err := tmpl.ExecuteTemplate(w, "contacts.page.tmpl", nil)
    if err != nil {
        log.Fatal(err)
    }
}
