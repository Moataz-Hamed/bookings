package main

import (
	"net/http"

	"github.com/Moataz-Hamed/bookings/pkg/config"
	"github.com/Moataz-Hamed/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	/* Basically when panic happens the application gets into panic situation which like saying iam dying and then dies
	Recoverer helps in this situation so when a panic happens the recoverer tells you of the problem and how to maybe
	solve it and keep the application alive
	*/
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
