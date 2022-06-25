package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Book     string `json:"book"`
	Author   string `json:"string"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Book: "First Book", Author: "First Author", Quantity: 1},
	{ID: "2", Book: "Second Book", Author: "Second Author", Quantity: 2},
}

// git.Context stores all the information related to the request, like headers, payload etc
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}

/* we need to retun pointer of type book otherwise we cannot return nil
since
nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type */

/* this won't work
func findBookById(id string) (book, error) {
	for _, b := range books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, errors.New("Book not found")
} */

func findBookById(id string) (*book, error) {
	for _, b := range books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, errors.New("Book not found for that particular ID")
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := findBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, *book)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.Run("localhost:8080")
}
