package main

import (
	"github.com/HironixRotifer/golang-hyperledger/app/pkg/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(cors.Default())

	app.POST("/book/create", controllers.CreateBook)
	app.POST("/book/create", controllers.DeleteBook)
	app.GET("/book/create", controllers.GetBookById)
	app.GET("/book/create", controllers.GetBooks)

	app.Run(":8080")
}
