package main

import (
	"fmt"
	"net/http"

	"github.com/e-razboinikov/learning-go/3.9/pkg/config"
	"github.com/e-razboinikov/learning-go/3.9/pkg/handlers"
	"github.com/e-razboinikov/learning-go/3.9/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Printf("Error while getting template canche, error: %s", err)
	}

	app.UseCache = false
	app.TemplateCache = templateCache

	repo := handlers.NewRepo(&app)
	handlers.NewHandles(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Application has started on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
