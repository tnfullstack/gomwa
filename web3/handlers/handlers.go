package handlers

import (
	"net/http"

	"github.com/tvn9/gomwa/web3/config"
	"github.com/tvn9/gomwa/web3/models"
	"github.com/tvn9/gomwa/web3/renders"
)

// Set Repo variable
var Repo *Repository

// Repository type struct
type Repository struct {
	AppCfg *config.AppConfig
}

// NewRepo creates new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		AppCfg: a,
	}
}

// NewHandler
func NewHandler(r *Repository) {
	Repo = r
}

// Home
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// passing data from to template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"
	renders.RenderTemplates(w, "home.html", &models.TemplateData{
		StrMap: stringMap,
	})
}

// About
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplates(w, "about.html", &models.TemplateData{})
}
