package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joe16723/booking/pkg/config"
	"github.com/joe16723/booking/pkg/handlers"
	"net/http"
)

// routes create a routes function from external lib
func routes(app *config.AppConfig) http.Handler {
	// routes with pat lib
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// routes with chi lib
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
