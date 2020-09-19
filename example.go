package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})
	router.GET("/greet", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Guest"
		}
		c.String(http.StatusOK, "Hello, %s!", name)
	})
	router.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	router.Run()
}
