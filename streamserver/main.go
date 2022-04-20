package main

import (
	"github.com/gin-gonic/gin"
)

type middleWareHandler struct {
	gin.ResponseWriter
	l *ConnLimiter
}

var (
	m = &middleWareHandler{
		l: NewConnLimiter(3),
	}
)

func checkLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !m.l.GetConn() {
			c.Writer.WriteString( "Too many requests.")
			return
		}
		defer m.l.ReleaseConn()
	}
}

func main(){
	//RegisterHandlers()
	r := gin.Default()

	r.Use(checkLimiter())
	r.Use(logResponseBody)

	r.LoadHTMLFiles("./testpage.html")

	r.GET("/videos/:vid-id", streamHandler())

	r.POST("/upload/:vid-id", uploadHandlers())

	r.GET("/testpage", testPageHandler())

	r.Run(":9000")
}
