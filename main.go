package main

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
	Mongodb_Repo.findMultipleBooks(10) // 10 is the limit

	// delete a single book
	Mongodb_Repo.deleteSingleBook(1)

	// update a book
	Mongodb_Repo.UpdateBookTitle(3, "The best num 3 book")
	Mongodb_Repo.increaseBookPrice(5, 9999)

	// find multiple books
	Mongodb_Repo.findMultipleBooks(10) // 10 is the limit

	// close CONNECTION
	Mongodb_Repo.disconnectDB()
}
