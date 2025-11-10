package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

type Controller struct {
	lib    services.LibraryManager
	reader *bufio.Reader
}

func NewController(lib services.LibraryManager) *Controller {
	return &Controller{
		lib:    lib,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *Controller) Run() {
	for {
		fmt.Println()
		fmt.Println("Library Management - Choose an option:")
		fmt.Println("1) Add book")
		fmt.Println("2) Remove book")
		fmt.Println("3) Borrow book")
		fmt.Println("4) Return book")
		fmt.Println("5) List available books")
		fmt.Println("6) List borrowed books by member")
		fmt.Println("7) Exit")
		choice := c.readInt("Enter choice: ")

		switch choice {
		case 1:
			c.handleAddBook()
		case 2:
			c.handleRemoveBook()
		case 3:
			c.handleBorrowBook()
		case 4:
			c.handleReturnBook()
		case 5:
			c.handleListAvailable()
		case 6:
			c.handleListBorrowed()
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func (c *Controller) handleAddBook() {
	id := c.readInt("Book ID: ")
	title := c.readString("Title: ")
	author := c.readString("Author: ")
	c.lib.AddBook(models.Book{ID: id, Title: title, Author: author})
	fmt.Println("Book added.")
}

func (c *Controller) handleRemoveBook() {
	id := c.readInt("Book ID to remove: ")
	c.lib.RemoveBook(id)
	fmt.Println("Book removed (if it existed).")
}

func (c *Controller) handleBorrowBook() {
	bookID := c.readInt("Book ID to borrow: ")
	memberID := c.readInt("Member ID: ")
	if err := c.lib.BorrowBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Book borrowed.")
}

func (c *Controller) handleReturnBook() {
	bookID := c.readInt("Book ID to return: ")
	memberID := c.readInt("Member ID: ")
	if err := c.lib.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Book returned.")
}

func (c *Controller) handleListAvailable() {
	books := c.lib.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	fmt.Println("Available books:")
	for _, b := range books {
		fmt.Printf("- [%d] %s by %s\n", b.ID, b.Title, b.Author)
	}
}

func (c *Controller) handleListBorrowed() {
	memberID := c.readInt("Member ID: ")
	books := c.lib.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books for this member or member not found.")
		return
	}
	fmt.Printf("Borrowed books for member %d:\n", memberID)
	for _, b := range books {
		fmt.Printf("- [%d] %s by %s\n", b.ID, b.Title, b.Author)
	}
}

func (c *Controller) readString(prompt string) string {
	for {
		fmt.Print(prompt)
		text, _ := c.reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
		fmt.Println("Input cannot be empty.")
	}
}

func (c *Controller) readInt(prompt string) int {
	for {
		fmt.Print(prompt)
		text, _ := c.reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if n, err := strconv.Atoi(text); err == nil {
			return n
		}
		fmt.Println("Please enter a valid number.")
	}
}


