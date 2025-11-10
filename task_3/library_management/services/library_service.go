package services

import (
	"errors"
	"sort"

	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	lib := &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
	lib.members[1] = models.Member{ID: 1, Name: "Alice"}
	lib.members[2] = models.Member{ID: 2, Name: "Bob"}
	return lib
}

func (l *Library) AddBook(book models.Book) {
	if book.Status == "" {
		book.Status = models.StatusAvailable
	}
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, mok := l.members[memberID]
	if !mok {
		return errors.New("member not found")
	}
	if book.Status == models.StatusBorrowed {
		return errors.New("book already borrowed")
	}
	book.Status = models.StatusBorrowed
	l.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	member, mok := l.members[memberID]
	if !mok {
		return errors.New("member not found")
	}
	idx := -1
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("book is not borrowed by this member")
	}
	member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[idx+1:]...)
	l.members[memberID] = member
	book.Status = models.StatusAvailable
	l.books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var result []models.Book
	for _, b := range l.books {
		if b.Status == models.StatusAvailable {
			result = append(result, b)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.members[memberID]
	if !ok {
		return []models.Book{}
	}
	result := make([]models.Book, len(member.BorrowedBooks))
	copy(result, member.BorrowedBooks)
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}
