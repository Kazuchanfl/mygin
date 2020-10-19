package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	router.Static("/a''", "./assets")
	// router.GET("/acount /:name''", acount)
	router.Run()
}
