package models

const (
	StatusAvailable = "Available"
	StatusBorrowed  = "Borrowed"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}
