package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type declaration example
type book struct {
	ID		string	`json:"id"`
	Title	string 	`json:"title"`
	Author	string	`json:"author"`
	Price	float64	`json:"price"`
}

// mocking data
var books = []book{
	{ID: "1", Title: "Aperte o F5: A transformação da Microsoft e a busca de um futuro melhor para todos", Author: "Satya Nadella", Price: 43.58},
	{ID: "2", Title: "The Comedians: Drunks, Thieves, Scoundrels and the History of American Comedy", Author: "Kliph Nesteroff", Price: 75.90},
	{ID: "3", Title: "The New Rules of Coffee: A Modern Guide for Everyone", Author: "Jordan Michelman & Zachary Carlsen", Price: 64.84},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.DELETE("/books/:id", removeBookByID)
	router.POST("/books", addBook)

	router.Run("localhost:3500")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func removeBookByID(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.ID == id {
			// Using append function to combine two slices
			// first slice is the slice of all the elements before the given index
			// second slice is the slice of all the elements after the given index
			// append function appends the second slice to the end of the first slice
			// returning a slice, so we store it in the form of a slice
			books = append(books[:index], books[index+1:]...)

			c.IndentedJSON(http.StatusOK, gin.H{"message": "book removed"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
