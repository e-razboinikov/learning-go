package handlers

import (
	"net/http"

	"github.com/e-razboinikov/learning-go/3.8/pkg/render"
)

/// Home is a home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

/// About is a about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
