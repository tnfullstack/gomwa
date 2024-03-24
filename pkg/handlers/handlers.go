package handlers

import (
	"net/http"

	"github.com/tvn9/gomwa/pkg/config"
	"github.com/tvn9/gomwa/pkg/models"
	"github.com/tvn9/gomwa/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.html", &models.AppData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, there!, I am here!"
	render.RenderTemplate(w, "about.html", &models.AppData{StringMap: stringMap})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.html", &models.AppData{})
}
