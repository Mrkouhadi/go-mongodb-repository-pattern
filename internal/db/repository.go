package db

import (
	"context"

	"github.com/mrkouhadi/go-graphql-mongo/internal/models"
)

type Repository interface {
	addBook(ctx context.Context, b models.Book) error
	addMultipleBooks(ctx context.Context, books []models.Book) error
	findeSingleBook(ctx context.Context, id int) (models.Book, error)
	findMultipleBooks(ctx context.Context, limit int64) ([]*models.Book, error)
	deleteSingleBook(ctx context.Context, id int) error
	deleteAllBooks(ctx context.Context) error
	UpdateBookTitle(ctx context.Context, id int, title string) error
	increaseBookPrice(ctx context.Context, id, inc int) error
	disconnectDB(ctx context.Context) error
}
