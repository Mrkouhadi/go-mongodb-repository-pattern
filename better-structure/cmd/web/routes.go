package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrkouhadi/go-graphql-mongo/internal"
	"github.com/mrkouhadi/go-graphql-mongo/internal/handlers"
)

func routes(app *internal.Config) http.Handler {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(LoadSession)

	// GET routes
	mux.Get("/book", handlers.Repo.Testt)
	mux.Get("/", handlers.Repo.Testhandler)

	// POST routes

	// render STATIC FILES
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// return our mux
	return mux
}
