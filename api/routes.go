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

	// GET routes
	mux.Get("/", handlers.Repo.Testhandler)

	// return our mux
	return mux
}
