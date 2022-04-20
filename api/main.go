package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()

	r.LoadHTMLGlob("tmpl/*")
	r.Static("/static", "./static")
	r.Static("/video", "./video")

	r.POST("/signup", func(c *gin.Context) {
		  
	})


	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", nil)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	r.GET("/course", func(c *gin.Context) {
		c.HTML(http.StatusOK, "course.html", nil)
	})

	r.GET("/course_video", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo.html", nil)
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/index-2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index-2.html", nil)
	})

	r.GET("/index-3", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index-3.html", nil)
	})

	r.GET("/no_event", func(c *gin.Context) {
		c.HTML(http.StatusOK, "no_event.html", nil)
	})

	r.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", nil)
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	r.GET("/single-blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "single-blog.html", nil)
	})

	r.GET("/single-course", func(c *gin.Context) {
		c.HTML(http.StatusOK, "single-course.html", nil)
	})

	r.GET("/single-event", func(c *gin.Context) {
		c.HTML(http.StatusOK, "single-event.html", nil)
	})

	r.GET("/single-team", func(c *gin.Context) {
		c.HTML(http.StatusOK, "single-team.html", nil)
	})

	r.GET("/team", func(c *gin.Context) {
		c.HTML(http.StatusOK, "team.html", nil)
	})



	r.GET("/mk", func(c *gin.Context) {
		c.HTML(http.StatusOK, "MK.html", nil)
	})

	r.Run(":2000")
}


