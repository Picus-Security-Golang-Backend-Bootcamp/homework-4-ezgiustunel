package library

import "errors"

var (
	ErrStockNotEnough = errors.New("there is not enough stock number for this book")
	ErrBookNotFound   = errors.New("book is not found")
)
