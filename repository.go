package main

import "context"

type Repository interface {
	addBook(ctx context.Context, b Book)
	addMultipleBooks(ctx context.Context, books []Book)
	findeSingleBook(ctx context.Context, id int)
	findMultipleBooks(ctx context.Context, limit int64)
	deleteSingleBook(ctx context.Context, id int)
	deleteAllBooks(ctx context.Context)
	UpdateBookTitle(ctx context.Context, id int, title string)
	increaseBookPrice(ctx context.Context, id, inc int)
	disconnectDB(ctx context.Context)
}
