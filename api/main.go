package main

import (
	"bysj_VEDIO/api/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Prepare() {
	session.LoadSessionsFromDB()
}

//type middleWareHandler struct {
//	c *gin.Context
//}
//
//func NewMiddleWareHandler(c *gin.Context) gin.HandlerFunc{
//		m := middleWareHandler{}
//		m.c = c
//		return m
//}
//
//func (m middleWareHandler) ServeHTTP(w gin.ResponseWriter, c *gin.Context) {
//	//check session
//	ValidateUserSession(c)
//
//	m.e.ServeHTTP(w, c.Request)
//}

func main(){
	Prepare()

	r := gin.Default()

	//r.Use(NewMiddleWareHandler(r))
	r.Use(logResponseBody)
	r.Use(ValidateUserSession)


	r.LoadHTMLGlob("tmpl/*")
	r.Static("/static", "./static")
	r.Static("/video", "./video")

	r.POST("/user/:username", Signin())

	r.POST("/user/", Signup())

	r.GET("/user/:username", GetUserInfo())

	r.POST("/user/:username/videos", AddNewVideo())

	r.GET("/user/:username/videos", ListAllVideos())

	r.DELETE("/user/:username/videos/:vid-id", DeleteVideo())

	r.POST("/video/:vid-id/comments", PostComment())

	//r.GET("/video/:vid-id/comments",ShowComments())





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


