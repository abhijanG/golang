package main

import "fmt"

type Book struct {
	name   string
	author string
	pages  int
}

type Shelf struct {
	position int
	book     Book
}

func structdemo() {
	book := Book{name: "Harry Potter", author: "J.K Rowling", pages: 800}
	fmt.Println(book)
	fmt.Println("Name :", book.name)
	fmt.Println("Author :", book.author)
	fmt.Println("Pages :", book.pages)
	fmt.Println()

	fmt.Println("Book Shelf")
	s := Shelf{position: 1, book: book}
	fmt.Println(s)
	fmt.Println("Book Position ", s.position)
	fmt.Println("Book Details ", s.book)
}
