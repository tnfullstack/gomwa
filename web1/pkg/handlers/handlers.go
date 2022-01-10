package handlers

import (
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
    tmpl.ExecuteTemplate(w, "index.tmpl", nil)
}

// About returns the web app about page
func About(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "about.tmpl", nil)
}

// Login returns the web app about page
func Login(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "login.tmpl", nil)
}

// Signup returns the web app about page
func Signup(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "signup.tmpl", nil)
}

// Contacts returns the web app about page
func Contacts(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "contacts.tmpl", nil)
}
