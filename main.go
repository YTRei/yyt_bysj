package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()
	r.LoadHTMLGlob("tmpl/*")
	r.Static("/static", "./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/index-2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index-2.html", nil)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	r.GET("/course", func(c *gin.Context) {
		c.HTML(http.StatusOK, "course.html", nil)
	})

	r.GET("/single-course", func(c *gin.Context) {
		c.HTML(http.StatusOK, "single-course.html", nil)
	})

	r.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", nil)
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	r.Run(":2000")
}