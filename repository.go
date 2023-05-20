package main

type Repository interface {
	addBook(b Book)
	addMultipleBooks(books []Book)
	findeSingleBook(id int)
	findMultipleBooks(limit int64)
	deleteSingleBook(id int)
	deleteAllBooks()
	disconnectDB()
}
