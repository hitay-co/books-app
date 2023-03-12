package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//  "errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	for _, book := range books {
		if book.ID == id {
			if book.Quantity <= 0 {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
				return
			}

			book.Quantity -= 1
			c.IndentedJSON(http.StatusOK, book)

			break
		}
	}

}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of books, looking for
	// an book whose ID value matches the parameter.
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	err := c.BindJSON(&newBook)

	if err != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)
	router.PUT("/checkout", checkoutBook)
	router.Run("localhost:8080")
}
