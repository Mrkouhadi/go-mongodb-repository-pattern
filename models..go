package main

type Book struct {
	ID     int
	TITLE  string
	AUTHOR string
	PRICE  float64
}

// dummy-data
var Book1 = Book{1, "Book number 1", "Kouhadi Bakr", 19.99}
var Book2 = Book{2, "Book number 2", "Bryan Aboubakr", 149.99}
var Book3 = Book{3, "Book number 3", "Omar Kouhadi", 399.99}
var Book4 = Book{4, "Book number 4", "Sheikh zoubir Kouhadi", 99.99}
var Book5 = Book{5, "Book number 5", "Mark Twain", 49.99}
