package main

import (
	"context"
	"time"
)

// main
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	Mongodb_Repo := CreateNewMongoDbRepo("mongodb://localhost:27017", ctx)
	///**The functions below are supposed to be called in handlers functions**/////
	// delete all documents already in collection

	Mongodb_Repo.deleteAllBooks(ctx)

	// addd a book
	Mongodb_Repo.addBook(ctx, Book1)

	// add multiple books
	Mongodb_Repo.addMultipleBooks(ctx, Book2, Book3, Book4, Book5) // can as many as we can here

	// find a book
	Mongodb_Repo.findeSingleBook(ctx, 2) // 2 is the id of the book

	// find multiple books
	Mongodb_Repo.findMultipleBooks(ctx, 10) // 10 is the limit

	// delete a single book
	Mongodb_Repo.deleteSingleBook(ctx, 1)

	// update a book
	Mongodb_Repo.UpdateBookTitle(ctx, 3, "The best num 3 book")
	Mongodb_Repo.increaseBookPrice(ctx, 5, 9999)

	// find multiple books
	Mongodb_Repo.findMultipleBooks(ctx, 10) // 10 is the limit

	// close CONNECTION
	Mongodb_Repo.disconnectDB(ctx)

}
