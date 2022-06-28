package handlers

import (
	"github.com/tvn9/gomwa/helloworld/models"
	"github.com/tvn9/gomwa/helloworld/pkg/config"
	"github.com/tvn9/gomwa/helloworld/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RendTemplate(w, "home.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["test"] = "Hello, again."
	render.RendTemplate(w, "about.tmpl", &models.TemplateData{
		StrMap: strMap,
	})
}
