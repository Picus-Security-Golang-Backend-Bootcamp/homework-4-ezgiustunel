package helper

import (
	"encoding/csv"
	"os"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-ezgiustunel/service/domain/book"
)

// ReadCsv: reads csv file
func ReadBookCsv(filename string) ([]book.Book, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []book.Book

	for _, line := range lines[1:] {
		id, _ := ConvertStringToInt(line[0])
		stockNumber, _ := ConvertStringToInt(line[2])
		pageNumber, _ := ConvertStringToInt(line[3])
		price, _ := ConvertStringToFloat64(line[4])
		authorId, _ := ConvertStringToInt(line[7])

		data := book.Book{
			ID:          id,
			Name:        line[1],
			StockNumber: stockNumber,
			PageNumber:  pageNumber,
			Price:       price,
			StockCode:   line[5],
			Isbn:        line[6],
			AuthorID:    authorId,
		}
		result = append(result, data)
	}

	return result, nil
}

func ReadAuthorCsv(filename string) ([]author.Author, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []author.Author

	for _, line := range lines[1:] {
		id, _ := ConvertStringToInt(line[0])

		data := author.Author{
			ID:         id,
			AuthorName: line[1],
		}
		result = append(result, data)
	}

	return result, nil
}
