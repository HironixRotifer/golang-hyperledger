package controllers

import (
	"net/http"

	"github.com/HironixRotifer/golang-hyperledger/app/pkg/gateway"
	"github.com/HironixRotifer/golang-hyperledger/app/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	title := c.Request.FormValue("title")
	author := c.Request.FormValue("author")
	publisher := c.Request.FormValue("publisher")

	_, err := gateway.Contract.SubmitTransaction("CreateBook", title, author, publisher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to submit transaction",
			"data":    err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction committed successfully",
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Request.FormValue("id")

	_, err := gateway.Contract.SubmitTransaction("DeleteBook", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to submit transaction",
			"data":    err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction committed successfully",
	})
}

func GetBookById(c *gin.Context) {
	id := c.Request.FormValue("id")

	evaluateResult, err := gateway.Contract.EvaluateTransaction("GetBookById", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to submit transaction",
			"data":    err,
		})
	}
	result := utils.FormatJSON(evaluateResult)

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction committed successfully",
		"data":    result,
	})
}

func GetBooks(c *gin.Context) {
	evaluateResult, err := gateway.Contract.EvaluateTransaction("GetBooks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to submit transaction",
			"data":    err,
		})
	}
	result := utils.FormatJSON(evaluateResult)

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction committed successfully",
		"data":    result,
	})
}
