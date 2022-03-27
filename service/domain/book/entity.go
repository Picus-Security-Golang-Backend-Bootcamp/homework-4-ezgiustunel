package book

import (
	errors "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/internal/library"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int `gorm:"primaryKey;autoIncrement" json:"Id"`
	StockNumber int
	PageNumber  int
	Price       float64
	Name        string
	StockCode   string
	Isbn        string
	AuthorID    int
	Author      author.Author `gorm:"foreignKey:AuthorID"`
}

// DecreaseStockNumber: checks and decreases stock number for the given book
func (b *Book) DecreaseStockNumber(bookNumber int) (*Book, error) {
	if b.StockNumber >= bookNumber {
		b.StockNumber -= bookNumber
		return b, nil
	}

	return nil, errors.ErrStockNotEnough
}
