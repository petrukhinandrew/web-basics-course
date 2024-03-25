package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/submit/:id", GetSubmitById)
	router.POST("/submit", AddSubmit)
	router.Run("mkn.edu:8080")
}
