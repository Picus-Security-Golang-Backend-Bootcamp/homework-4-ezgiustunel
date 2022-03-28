package network

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/book"
	"github.com/gorilla/mux"
)

var BookList []book.Book
var AuthorList []author.Author
var books []book.Book
var authors []author.Author

var BookRepository *book.BookRepository
var AuthorRepository *author.AuthorRepository

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books = BookRepository.FindAll()
	json.NewEncoder(w).Encode(books)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors = AuthorRepository.FindAll()
	json.NewEncoder(w).Encode(authors)
}

func GetBooksWithParam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	searchText := params["searchText"]
	books, err := BookRepository.FindByBookName(searchText)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["authorId"])

	var book book.Book
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = rand.Intn(10000000)
	book.AuthorID = id

	BookRepository.InsertData(book)
	json.NewEncoder(w).Encode(book)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author author.Author

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&author)
	author.ID = rand.Intn(10000000)

	AuthorRepository.InsertData(author)
	json.NewEncoder(w).Encode(author)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	book := BookRepository.DeleteById(id)

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	book, err := BookRepository.FindById(id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bookNumber, _ := strconv.Atoi(params["bookNumber"])
	_, errStock := book.DecreaseStockNumber(bookNumber)

	if errStock != nil {
		fmt.Println(errStock.Error())
		return
	}
	BookRepository.Update(book)
}
