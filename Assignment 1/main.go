package main

import (
	"github.com/DenisLi/Assignment1/Library"
)

func main() {
	//library test
	// fmt.Println("salam")
	library := new(Library.Library)
	// book1 := Library.Book{

	// 	ID:         1,
	// 	Title:      "Преступление и Наказание",
	// 	Author:     "Достоевский",
	// 	IsBorrowed: false,
	// }
	// book2 := Library.Book{

	// 	ID:         2,
	// 	Title:      "Идиот",
	// 	Author:     "Достоевский",
	// 	IsBorrowed: false,
	// }
	// book3 := Library.Book{

	// 	ID:         3,
	// 	Title:      "Братья Карамазовы",
	// 	Author:     "Достоевский",
	// 	IsBorrowed: true,
	// }
	// library.AddBook(book1)
	// library.AddBook(book2)
	// library.AddBook(book3)
	// library.BorrowBook(book3)
	// library.BorrowBook(book1)
	// library.ListAvailableBooks()
	// library.ReturnBook(book1)
	// library.ReturnBook(book3)
	// library.ReturnBook(book2)
	// library.ListAvailableBooks()
	// fmt.Println("Welcome to library!\n\nCommands:\n1. Add a book\n2. Borrow a book\n3. Return a book\n4. List all available books")
	library.CLI()
}
