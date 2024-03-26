package main

import (
	_ "embed"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/resp-template.tmpl")
	router.MaxMultipartMemory = 8 << 20
	router.Use(static.Serve("/", static.LocalFile("../www/", false)))
	router.GET("/api/submit/:id", GetSubmitById)
	router.POST("/api/submit", AddSubmit)

	router.RunTLS("mkn.edu:443", "../mkn.edu.crt", "../mkn.edu.key")
}
