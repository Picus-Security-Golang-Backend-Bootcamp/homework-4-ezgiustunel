package helper

import (
	"fmt"
	"strconv"
)

// Method for converting string input to int
// if it is a number then return the number, otherwise return -1
func ConvertStringToInt(input string) (int, error) {
	result, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func ConvertStringToFloat64(input string) (float64, error) {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return -1, err
	}
	return result, nil
}

// Method for print messages to user in the console.
func PrintMessagesToConsole() {
	fmt.Printf("\n--Invalid Input--\n\n")
	fmt.Println("You can use the methods below to make some actions on book list")
	fmt.Println("list: Lists the books")
	fmt.Println("search \"bookname\": searches the bookname given in the book list")
	fmt.Println("buy: you can buy books")
	fmt.Printf("delete: you can delete a book from book list\n\n")
}
