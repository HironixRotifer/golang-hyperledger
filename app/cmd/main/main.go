package main

import (
	"fmt"

	"github.com/HironixRotifer/golang-hyperledger/app/pkg/controllers"
	"github.com/HironixRotifer/golang-hyperledger/app/pkg/gateway"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(cors.Default())
	gateway.Init()

	app.POST("/book/create", controllers.CreateBook)
	app.POST("/book/delete", controllers.DeleteBook)
	app.GET("/book/getBook/{id}", controllers.GetBookById)
	app.GET("/book/getBooks", controllers.GetBooks)

	fmt.Println(gateway.Contract)

	app.Run(":8080")
}
