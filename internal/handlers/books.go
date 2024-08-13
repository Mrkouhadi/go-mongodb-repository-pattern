package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mrkouhadi/go-graphql-mongo/internal/utils"
)

func (handlerRepo *Repository) Testhandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// find a single book
	book, err := handlerRepo.DB.FindSingleBook(ctx, 2) // 2 is the id of the book
	if err != nil {
		utils.HandleDBError("finding a single book: ", err)
	}
	log.Println(book)
	/*
		// The rest of methods can be used like this:
			// delete all book
			if err := handlerRepo.DB.DeleteAllBooks(ctx); err != nil {
				handleDBError("Deleting all books: ", err)
			}
			err = handlerRepo.DB.DeleteAllBooks(ctx)
			err = handlerRepo.DB.AddBook(ctx, models.Book1) // ctx, new book
			err = handlerRepo.DB.AddMultipleBooks(ctx, models.Book2, models.Book3, models.Book4, models.Book5)
			_, err = handlerRepo.DB.FindMultipleBooks(ctx, 10)                  // 10 is the limit
			err = handlerRepo.DB.DeleteSingleBook(ctx, 1)                       // ctx,id
			err = handlerRepo.DB.UpdateBookTitle(ctx, 3, "The best num 3 book") // ctx, new title
			err = handlerRepo.DB.IncreaseBookPrice(ctx, 5, 9999)                // ctx,id,newPrice
			err = handlerRepo.DB.DisconnectDB(ctx)

			if err != nil {
				handleDBError("The rest of methods: ", err)
			}
	*/
}
