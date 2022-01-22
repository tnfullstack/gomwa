package main

import (
	"log"
	"net/http"

	"github.com/tvn9/gomwa/web3/config"
	"github.com/tvn9/gomwa/web3/handlers"
	"github.com/tvn9/gomwa/web3/renders"
)

func main() {

	var appCfg config.AppConfig

	tc, err := renders.CreateTmplCache()
	if err != nil {
		log.Fatal(err)
	}

	appCfg.TmplCache = tc

	appCfg.UseCache = true

	repo := handlers.NewRepo(&appCfg)
	handlers.NewHandler(repo)

	renders.NewTemplates(&appCfg)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	srv := http.Server{
		Addr: ":8080",
	}

	srv.ListenAndServe()
}
