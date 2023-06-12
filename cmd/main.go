package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tvn9/gomwa/pkg/config"
	"github.com/tvn9/gomwa/pkg/handlers"
	"github.com/tvn9/gomwa/pkg/render"
)

const portNum = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Println("Starting server on port:", portNum)
	_ = http.ListenAndServe(portNum, nil)
}
