package main

import (
	"log"
	"net/http"

	"github.com/tvn9/gomwa/pkg/config"
	"github.com/tvn9/gomwa/pkg/handlers"
	"github.com/tvn9/gomwa/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("connot create template cache. %v\n", err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/contact", handlers.Repo.Contact)

	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("fail to start server %v\n", err)
	}
}
