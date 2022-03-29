package handlers

import (
	"net/http"

	"github.com/e-razboinikov/learning-go/3.9/pkg/config"
	"github.com/e-razboinikov/learning-go/3.9/pkg/render"
)

/// Repo - the repository used by the handlers
var Repo *Repository

/// Repository is a repository type
type Repository struct {
	App *config.AppConfig
}

/// NewRepo creates a new repository
func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

/// NewHandlers sets the repository for the handlers
func NewHandles(r *Repository) {
	Repo = r
}

/// Home is a home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

/// About is a about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
