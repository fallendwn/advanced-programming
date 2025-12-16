package Library

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var id int = 1

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (library *Library) AddBook(book Book) {

	if library.Books == nil {
		library.Books = make(map[string]Book)
	}
	library.Books[book.ID] = book
	fmt.Printf("Successfully add book with name %s\n", library.Books[book.ID].Title)
}

func (library *Library) BorrowBook(book Book) {
	var thisBook Book = library.Books[book.ID]
	if !thisBook.IsBorrowed {
		thisBook.IsBorrowed = true
		fmt.Printf("You successfully borrowed a book %s\n", thisBook.Title)
		library.Books[book.ID] = thisBook

	} else {
		fmt.Printf("Book %s is already borrowed!\n", thisBook.Title)
	}

}

func (library *Library) ReturnBook(book Book) {
	var thisBook Book = library.Books[book.ID]
	if thisBook.IsBorrowed {
		thisBook.IsBorrowed = false
		fmt.Printf("You successfully returned a book named %s\n", thisBook.Title)
		library.Books[book.ID] = thisBook
	} else {
		fmt.Printf("Book is on its place, no need to return.\n")
	}

}

func (library *Library) ListAvailableBooks() {
	var counter int = 0
	for _, book := range library.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
			counter++
		}
	}
}

func (library *Library) CLI() {
	var scanner = bufio.NewScanner(os.Stdin)
	var cursor string
	for {
		fmt.Printf("\nWelcome to library!\n\nCommands:\n1. Add a book\n2. Borrow a book\n3. Return a book\n4. List all available books\n5. Exit\n\n")
		scanner.Scan()
		cursor = scanner.Text()
		switch cursor {
		case "1":
			fmt.Println("Enter book title")
			scanner.Scan()
			title := scanner.Text()
			fmt.Println("Enter book author")
			scanner.Scan()
			author := scanner.Text()
			newBook := Book{
				ID:         strconv.Itoa(id),
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}
			id += 1
			library.AddBook(newBook)
		case "2":
			var bookId string
			fmt.Println("Enter an ID of book you want to borrow!")
			library.ListAvailableBooks()
			scanner.Scan()
			bookId = scanner.Text()
			book, exists := library.Books[bookId]
			if exists {
				library.BorrowBook(book)
			} else {
				fmt.Println("There is no such book.")
			}
		case "3":
			var bookId string
			fmt.Println("Enter an ID of book you want to return!")
			scanner.Scan()
			bookId = scanner.Text()
			book, exists := library.Books[bookId]
			if exists {
				library.ReturnBook(book)
			} else {
				fmt.Println("There is no such book.")
			}
		case "4":
			library.ListAvailableBooks()
		case "5":
			fmt.Println("Goodbye!")
			return
		}
	}
}
