package main

// import "github.com/gin-gonic/gin"
import (
	"fmt"

	"github.com/jinzhu/gorm"
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

	db, _ := gorm.Open("sqlite3", "C:\\SQLiteStudio\\books.db")
	defer db.Close()
	db.AutoMigrate(&Book{})

	b1 := Book{ID: 1, Title: "Labyrinth", Author: "Kate Moose"}
	db.Create(&b1)
	fmt.Println(b1.Title)

}
