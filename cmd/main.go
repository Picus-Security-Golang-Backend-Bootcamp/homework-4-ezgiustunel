package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/network"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/infrastructure"
	"github.com/gorilla/mux"
)

type Book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	db := infrastructure.ConnectDB("postgres://ezgiustunel:pass@localhost:5432/library")
	network.AuthorRepository = author.NewAuthorRepository(db)
	network.BookRepository = book.NewBookRepository(db)

	network.AuthorRepository.Migration()
	network.BookRepository.Migration()

	network.BookList, _ = helper.ReadBookCsv("book.csv")
	network.AuthorList, _ = helper.ReadAuthorCsv("author.csv")

	for _, book := range network.BookList {
		network.BookRepository.InsertData(book)
	}

	for _, author := range network.AuthorList {
		network.AuthorRepository.InsertData(author)
	}
}

func main() {
	r := mux.NewRouter()
	fmt.Println(r)

	r.HandleFunc("/api/books", network.GetBooks).Methods("GET")
	r.HandleFunc("/api/authors", network.GetAuthors).Methods("GET")
	r.HandleFunc("/api/books/{searchText}", network.GetBooksWithParam).Methods("GET")
	r.HandleFunc("/api/books/{authorId}", network.CreateBook).Methods("POST")
	r.HandleFunc("/api/authors", network.CreateAuthor).Methods("POST")
	r.HandleFunc("/api/books/{id}", network.DeleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}/{bookNumber}", network.UpdateBook).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
