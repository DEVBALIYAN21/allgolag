package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Book represents the structure of a book in the library
type Book struct {
	ID     int    `json:"id"`      // Book ID
	Title  string `json:"title"`   // Book Title
	Author string `json:"author"`  // Book Author
}

// In-memory slice to hold books
var books = []Book{
	{ID: 1, Title: "Book 1", Author: "Auth 1"},
	{ID: 2, Title: "Book 2", Author: "Auth 2"},
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Routes
	r.GET("/books", getBooks)         // Get all books
	r.POST("/books", addBook)         // Add a new book
	r.DELETE("/books/:id", deleteBook) // Delete a book by ID

	// Start the server on port 8080
	fmt.Println("Library API running on port 8080")
	r.Run(":8080")
}

// Handler to get the list of books
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// Handler to add a new book
func addBook(c *gin.Context) {
	var newBook Book

	// Bind the incoming JSON to the newBook struct
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new ID for the book (ID is auto-incremented)
	newBook.ID = len(books) + 1

	// Add the new book to the in-memory slice
	books = append(books, newBook)

	// Return the added book in the response
	c.JSON(http.StatusCreated, newBook)
}

// Handler to delete a book by ID
func deleteBook(c *gin.Context) {
	// Get the book ID from the URL parameter
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Find the book  to delete

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	// If book is not found
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
