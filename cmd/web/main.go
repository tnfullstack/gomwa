package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tvn9/gomwa/pkg/config"
	"github.com/tvn9/gomwa/pkg/handlers"
	"github.com/tvn9/gomwa/pkg/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("connot create template cache. %v\n", err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Starting server on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	log.Fatal(srv.ListenAndServe())
}
