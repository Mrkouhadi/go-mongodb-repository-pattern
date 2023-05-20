package main

type Repository interface {
	addBook(b Book)
	addMultipleBooks(books []Book)
	findeSingleBook(id int)
	findMultipleBooks(limit int64)
	deleteSingleBook(id int)
	deleteAllBooks()
	UpdateBookTitle(id int, title string)
	increaseBookPrice(id, inc int)
	disconnectDB()
}
