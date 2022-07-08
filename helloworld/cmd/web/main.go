package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tvn9/gomwa/helloworld/pkg/config"
	"github.com/tvn9/gomwa/helloworld/pkg/handlers"
	"github.com/tvn9/gomwa/helloworld/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNum = ":8080"

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
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting server on port %s\n", portNum)
	//log.Fatal(http.ListenAndServe(portNum, nil))
	log.Fatal(srv.ListenAndServe())
}
