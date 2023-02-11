package main

import (
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3307)/go_book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&book.Book{})
	bookRepository := book.AssignRepository(db)
	bookService := book.AssignService(bookRepository)
	bookHandler := book.AssignHandler(bookService)

	// bookItems, err := bookService.FindAll()
	// for _, book := range bookItems {
	// 	fmt.Println("Books :", book)
	// }

	// bookItem, err := bookService.FindById(2)
	// // if err != nil {}
	// fmt.Println("Book :", bookItem)

	// bookModel := book.BookInputType{
	// 	Title:       "Trade Thinking Analytic",
	// 	SubTitle:    "About how to self development improve bad habbits to good habbits",
	// 	Price:       56000,
	// 	Description: "Ut excepteur cupidatat ut aliquip nulla fugiat cupidatat esse. Reprehenderit mollit dolore est laboris nisi. Lorem incididunt id aliquip occaecat cupidatat fugiat do esse fugiat laboris. Deserunt aute duis minim in incididunt non. Aliquip nostrud incididunt eu proident ea. Nisi qui occaecat reprehenderit enim et deserunt enim ex.",
	// }
	// newBookItem, err := bookService.Create(bookModel)
	// fmt.Println("Book Create :", newBookItem)

	// Create
	// bookModel := book.Book{}
	// bookModel.Title = "Anak kambing saya"
	// bookModel.SubTitle = "About how to self development improve bad habbits to good habbits"
	// bookModel.Price = 56000
	// bookModel.Discount = 0
	// bookModel.Rating = 2
	// bookModel.Description = "Ut excepteur cupidatat ut aliquip nulla fugiat cupidatat esse. Reprehenderit mollit dolore est laboris nisi. Lorem incididunt id aliquip occaecat cupidatat fugiat do esse fugiat laboris. Deserunt aute duis minim in incididunt non. Aliquip nostrud incididunt eu proident ea. Nisi qui occaecat reprehenderit enim et deserunt enim ex."

	// err = db.Create(&bookModel).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error creating book")
	// 	fmt.Println("========================")
	// }

	// fmt.Println("Success create book :", bookModel)

	// Read
	// var bookItem book.Book
	// // err = db.Debug().First(&bookItem).Error
	// err = db.Debug().First(&bookItem, 2).Error
	// // err = db.Debug().Last(&bookItem).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error finding book")
	// 	fmt.Println("========================")
	// }
	// fmt.Println("Title :", bookItem.Title)
	// fmt.Println("book :", bookItem)

	// var bookItems []book.Book
	// // err = db.Debug().Find(&bookItems).Error
	// err = db.Debug().Where("title LIKE ?", "%atomic%").Find(&bookItems).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error finding books")
	// 	fmt.Println("========================")
	// }

	// for _, v := range bookItems {
	// 	fmt.Println("========================")
	// 	fmt.Println("TItle :", v.Title)
	// 	fmt.Println("book :", v)
	// }

	// Update
	// var bookItem book.Book
	// err = db.Debug().Where("id = ?", 1).Find(&bookItem).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error finding book")
	// 	fmt.Println("========================")
	// }

	// bookItem.Title = "Man of Taichi (Revised Edition)"
	// err = db.Save(&bookItem).Error

	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error updating book")
	// 	fmt.Println("========================")
	// }

	// fmt.Println("Success update book :", bookItem)

	// Delete
	// var bookItem book.Book
	// err = db.Debug().Where("id = ?", 3).Find(&bookItem).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error finding book")
	// 	fmt.Println("========================")
	// }

	// err = db.Delete(&bookItem).Error
	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error deleting book")
	// 	fmt.Println("========================")
	// }

	// fmt.Println("Success delete book :", bookItem)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/ping", bookHandler.TestHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/books", bookHandler.BookHandler)
	v1.GET("/books/:id", bookHandler.DetailBookHandler)
	v1.POST("/books", bookHandler.AddBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	// router.Run() // default
	router.Run(":8888") // changed port

	// main
	// handler
	// service
	// repository
	// db orm
	// db mysql
}
