package main

import (
	"bookself-api/book"
	"bookself-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bookself-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connection Error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title : ", book.Title)
	// }
	
	// bookInput := book.BookInput{
	// 	Title : "Atomic Habbit",
	// 	Description: "Easy Read book",
	// 	Price: "95000",
	// 	Rating: 5,
	// }


	// bookService.Create(bookInput)


	//CRUD DATA


	//Create Method
	// book := book.Book{}
	// book.Title = "Do it Now"
	// book.Price = 190000
	// book.Rating = 5
	// book.Description = "Biarkan Diriku Sendiri Sekarang"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }


	//Read Method (Find)
	// var books []book.Book

	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }

	// for _, b := range books{
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("book object : %v", b)
	// }

	//Read Method (First)
	// var book book.Book

	// err = db.Debug().First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }
	// 	fmt.Println("Title : ", book.Title)
	// 	fmt.Println("book object : %v", book)


	//Read Method (Where and Find)
	// var books []book.Book

	// err = db.Debug().Where("Rating = ?", 5).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }

	// for _, b := range books{
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("book object : %v", b)
	// }

	//Update Method Save
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }

	// book.Title = "Palagan Nusantara 2"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }

	// var book book.Book
	// err = db.Debug().Where("Title = ?", "Do it Now").Find(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("Error Cuy")
	// }

	

	router := gin.Default()

	v1 := router.Group("/v1")


	v1.GET("/", bookHandler.RootHandler)
	router.GET("/GetAllBooks", bookHandler.GetAllBooks)
	router.POST("/addBooks", bookHandler.PostBooksHandler)
	router.GET("/GetBook/:id", bookHandler.GetBook)

	router.Run(":2000")
}





