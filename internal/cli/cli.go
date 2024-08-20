package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: books <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ...")
		}

		// Pegar todos os dados do apartir do segundo parametro
		bookIDs := os.Args[2:]
		cli.simulateReading(bookIDs)
	}

}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Error searching books:", err)
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Printf("%d books found\n", len(books))

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n",
			book.ID, book.Title, book.Author, book.Genre)
	}

}

func (cli *BookCLI) simulateReading(bookIDsStr []string) {
	var bookIds []int
	for _, idStr := range bookIDsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println()
			continue
		}
		bookIds = append(bookIds, id)
	}

	responses := cli.service.SimulateMultipleReadings(bookIds, 5*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}
