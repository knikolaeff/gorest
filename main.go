package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price float64 `json:"price"`
}

var books = []book {
	{"1", "Crime and Punishment", "Fedor Dostoevsky", 10.99},
	{"2", "War and Peace", "Leo Tolstoy", 25.99},
	{"3", "Sherlock Holmes","Arthur Conan Doyle", 12.99},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found!"})
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