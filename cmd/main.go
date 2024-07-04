package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("../web/templates/default.html")
	router.Static("/static", "../static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default.html", gin.H{
			"title": "Task list",
		})
	})

	router.Run(":8080")
}
