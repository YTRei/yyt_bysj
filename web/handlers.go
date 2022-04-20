package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"yyt/config"
)

type HomePage struct {
	Name string

}

type UserPage struct {
	Name string
}

type responseBodyWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWrite) WriteString (s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func logResponseBody(c *gin.Context) {
	w := &responseBodyWrite{
		ResponseWriter: c.Writer,
		body:           &bytes.Buffer{},
	}
	c.Writer = w
	c.Next()

	fmt.Println("Response body: " + w.body.String())
}

func homeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cname, err1 := c.Request.Cookie("username")

		sid, err2 := c.Request.Cookie("session")

		fmt.Println(cname, sid)

		if err1 != nil || err2 != nil {
			//p := &HomePage{Name: "jack"}
			c.HTML(http.StatusOK, "signin.html", nil)
			return
		}

		if len(cname.Value) != 0 && len(sid.Value) != 0 {
			c.Redirect(http.StatusFound, "/userhome")
			return
		}
	}
}

func userHomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cname, err1 := c.Request.Cookie("username")
		_, err2 := c.Request.Cookie("session")

		if err1 != nil || err2 != nil {
			c.Redirect(http.StatusFound, "/user")
			return
		}

		fname := c.Request.FormValue("username")

		var p *UserPage
		if len(cname.Value) != 0 {
			p = &UserPage{Name: cname.Value}
		}else if len(fname) != 0 {
			p = &UserPage{Name: fname}
		}

		c.HTML(http.StatusOK, "userhome.html", p)
	}
}

func apiHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method != http.MethodPost {
			c.Writer.WriteString("Api not recognized, bad request")
			return
		}

		res, _ := ioutil.ReadAll(c.Request.Body)
		apibody := &ApiBody{}
		if err := json.Unmarshal(res, apibody); err != nil {
			c.Writer.WriteString("request body is not correct")
			return
		}

		request(apibody, c, c.Writer)
		defer c.Request.Body.Close()
	}
}

func proxyVideoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := url.Parse("http://" + config.GetLbAddr() + ":9000/")
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func proxyUploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := url.Parse("http://" + config.GetLbAddr() + ":9000/")
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}





