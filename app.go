package main

import (
	"cloud-run-router/responses"
	"cloud-run-router/settings"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/go-chi/chi/v5"
)

type App struct {
	Router      *chi.Mux
	FirebaseApp *firebase.App
}

func (app *App) Initialize(config settings.Config) {
	app.Router = chi.NewRouter()

	app.Router.Get("/", app.getIndex)

	if config.ExposeRoutes {
		app.Router.Get("/_routes", func(w http.ResponseWriter, r *http.Request) {
			app.getRoutes(w, r, config)
		})
	}
}

func (app *App) Run(port string) {
	log.Printf("Listening on http://0.0.0.0%s", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) getIndex(w http.ResponseWriter, r *http.Request) {
	responses.Ok(w, map[string]string{"status": "up"})
}

func (app *App) getRoutes(w http.ResponseWriter, r *http.Request, config settings.Config) {
	responses.Ok(w, config.Routes)
}
