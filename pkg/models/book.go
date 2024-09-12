package models

import (
	"bookstore/pkg/config"
	"database/sql"
	"fmt"
)

var db *sql.DB

// book table
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Stock       int    `json:"Stock"`
}

// Initialize the database connection
func init() {
	config.Connect()
	db = config.GetDB()
}

// updates the stock every each order
func UpdateStock(id int) error {
	updateStockQuery := `UPDATE books SET Stock = Stock - 1 WHERE ID = ?`
	_, err := db.Exec(updateStockQuery, id)
	if err != nil {
		return fmt.Errorf("error updating book stock: %v", err)
	}
	return nil
}

// CreateBook inserts a new book into the database
func CreateBook(b *Book) (*Book, error) {
	id, err := MaxID()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}
	b.ID = int(id) + 1

	query := "INSERT INTO books (id,name, author, publication) VALUES (?,?, ?, ?)"
	_, errr := db.Exec(query, b.ID, b.Name, b.Author, b.Publication)
	if errr != nil {
		return nil, fmt.Errorf("failed to insert book: %v", errr)
	}

	return b, nil
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() ([]Book, error) {
	query := "SELECT id, name, author, publication,Stock FROM books"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %v", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publication, &book.Stock); err != nil {
			return nil, fmt.Errorf("failed to scan book: %v", err)
		}
		books = append(books, book)
	}

	return books, nil
}

// GetBookById retrieves a book by its ID
func GetBookById(id int) (*Book, error) {
	query := "SELECT id, name, author, publication,Stock FROM books WHERE id = ?"
	row := db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Name, &book.Author, &book.Publication, &book.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to get book: %v", err)
	}

	return &book, nil
}

// DeleteBook removes a book from the database by its ID
func DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}

	return nil
}

// UpdateBook updates the details of an existing book
func UpdateBook(b *Book) (*Book, error) {
	query := "UPDATE books SET name = ?, author = ?, publication = ? WHERE id = ?"
	_, err := db.Exec(query, b.Name, b.Author, b.Publication, b.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update book: %v", err)
	}

	return b, nil
}

// get the max id exist in the table
func MaxID() (int, error) {
	var maxID int
	query := "SELECT MAX(id) FROM books"
	row := db.QueryRow(query)
	err := row.Scan(&maxID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to retrieve max ID: %v", err)
	}
	return maxID, nil
}
