package main

// dummy-data
var Book1 = Book{1, "Book number 1", "Kouhadi Bakr", 19.99}
var Book2 = Book{2, "Book number 2", "Bryan Aboubakr", 149.99}
var Book3 = Book{3, "Book number 3", "Omar Kouhadi", 399.99}
var Book4 = Book{4, "Book number 4", "Sheikh zoubir Kouhadi", 99.99}
var Book5 = Book{5, "Book number 5", "Mark Twain", 49.99}

// main
func main() {
	Mongodb_Repo := CreateNewMongoDbRepo("mongodb://localhost:27017")

	// delete all documents already in collection
	Mongodb_Repo.deleteAllBooks()

	// addd a book
	Mongodb_Repo.addBook(Book1)

	// add multiple books
	Mongodb_Repo.addMultipleBooks(Book2, Book3, Book4, Book5 /*can as many as we can here */)

	// find a book
	Mongodb_Repo.findeSingleBook(2) // 2 is the id of the book

	// find multiple books
	Mongodb_Repo.findMultipleBooks(4) // 4 is the limit

	// delete a single book
	Mongodb_Repo.deleteSingleBook(1)

	// close CONNECTION
	Mongodb_Repo.disconnectDB()
}
