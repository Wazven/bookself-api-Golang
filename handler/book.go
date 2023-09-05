package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"bookself-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler)RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Aldi Rhiyadi",
		"Bio":  "A Back-End Engineer",
	})
}

func (handler *bookHandler) GetAllBooks(c *gin.Context){
	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin. H{
			"errors": err,
		})
		return
	}
	var booksResponse []book.BookResponse

	for _, b := range books{
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (handler *bookHandler) GetBook(c *gin.Context){
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	b, err := handler.bookService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin. H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)
	
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (handler *bookHandler) PostBooksHandler(c *gin.Context){
	//title, price
	var bookInput book.BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	
	}

	book, err := handler.bookService.Create(bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (handler *bookHandler) UpdateBook(c *gin.Context){
	//title, price
	var bookInput book.BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	book, err := handler.bookService.Update(id, bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (handler *bookHandler) DeleteBook(c *gin.Context){
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	b, err := handler.bookService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(b),
	})

}

func convertToBookResponse(b book.Book) book.BookResponse{
	return book.BookResponse{
			Title: b.Title,
			Price: b.Price,
			Description: b.Description,
			Rating: b.Rating,
			ID: b.ID,
	}
}