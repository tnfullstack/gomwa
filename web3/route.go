package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tvn9/gomwa/web3/config"
	"github.com/tvn9/gomwa/web3/handlers"
)

func routes(acf *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
