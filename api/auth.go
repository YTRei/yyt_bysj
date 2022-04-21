package main

import (
	"bysj_VEDIO/api/defs"
	"bysj_VEDIO/api/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"


func ValidateUserSession(c *gin.Context) {
	//fmt.Println("VVVVVVVVVVVVVV run!!!!!!!")

		sid := c.Request.Header.Get(HEADER_FILED_SESSION)

		if len(sid) == 0 {
			return
		}

		uname, ok := session.IsSessionExpired(sid)
		if ok {
			fmt.Println("unameeee : ", uname)
			return
		}
		c.Request.Header.Add(HEADER_FILED_UNAME, uname)
	ubody:= c.Request.Header.Clone()
	fmt.Println("UUUUUUUUUUUUUUUUUUubody: ",ubody)
}

func ValidateUser (w gin.ResponseWriter, r *http.Request ) bool {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	ubody:= r.Header.Clone()
	fmt.Println("ubody: ",ubody)
	if len(uname) == 0 {
		fmt.Println(uname)
		w.WriteString(defs.ErrorRequestBodyParseFailed.Error)
		return false
	}

	return true
}


