package handlers

import (
	"github.com/mrkouhadi/go-graphql-mongo/internal"
	"github.com/mrkouhadi/go-graphql-mongo/internal/db"
)

type Repository struct {
	App *internal.Config
	DB  db.MongodbRepo
}

var Repo *Repository

func NewRepo(app *internal.Config, db *db.MongodbRepo) *Repository {
	return &Repository{
		App: app,
		DB:  *db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
