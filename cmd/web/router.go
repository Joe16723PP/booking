package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joe16723/booking/pkg/config"
	"github.com/joe16723/booking/pkg/handlers"
	"net/http"
)

// routes create a routes function from external lib
func routes(_ *config.AppConfig) http.Handler {
	// routes with chi lib
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// serving static file eg. image
	fileServer := http.FileServer(http.Dir("./assets/"))
	mux.Handle("/assets/*", http.StripPrefix("/assets", fileServer))
	return mux
}
