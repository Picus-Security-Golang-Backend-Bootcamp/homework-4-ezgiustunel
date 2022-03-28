package book

import (
	"errors"
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/internal/library"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

// NewBookRepository: creates new book repo
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

// Migration: migrates model and db
func (b *BookRepository) Migration() {
	b.db.AutoMigrate(&Book{})
}

// InsertData: insert data to db
func (b *BookRepository) InsertData(book Book) {
	b.db.Where(Book{ID: book.ID}).Attrs(Book{StockNumber: book.StockNumber, PageNumber: book.PageNumber, Price: book.Price, Name: book.Name, StockCode: book.StockCode, Isbn: book.Isbn, AuthorID: book.AuthorID}).FirstOrCreate(&book)
}

// FindAll: finds all elements in db
func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)

	return books
}

// FindByBookName: find element by book name
func (b *BookRepository) FindByBookName(bookName string) ([]Book, error) {
	var books []Book
	b.db.Where("Name LIKE ?", "%"+bookName+"%").Find(&books)

	if books == nil {
		return nil, library.ErrBookNotFound
	}
	return books, nil
}

// FindById: find element by id
func (b *BookRepository) FindById(id int) (*Book, error) {
	var book *Book
	result := b.db.First(&book, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Book not found with id : %d", id)
		return nil, library.ErrBookNotFound
	}
	return book, nil
}

// Update: updates the database
func (b *BookRepository) Update(book *Book) error {
	result := b.db.Save(book)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("updated successfully.")
	return nil
}

// DeleteById: deletes element from database
func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("deleted successfully.")
	return nil
}
