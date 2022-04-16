package main

// import "github.com/gin-gonic/gin"
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })

	// r.Run()

	fmt.Println("Go ORM Tutorial")

	db, err := gorm.Open("sqlite3", "books.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Book{})

	b1 := Book{ID: 1, Title: "Labyrinth", Author: "Kate Moose"}
	db.Create(&b1)
	fmt.Println(b1.Title)

}
