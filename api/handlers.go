package main

import (
	"bysj_VEDIO/api/dbops"
	"bysj_VEDIO/api/defs"
	"bysj_VEDIO/api/session"
	"bysj_VEDIO/api/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
)

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

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("%s", res)
		ubody := &defs.UserCredential{}

		if err := json.Unmarshal(res, ubody); err != nil {
			c.Writer.WriteString(defs.ErrorRequestBodyParseFailed.Error)
			return
		}
		//fmt.Println(ubody.Age)
		age, err := strconv.Atoi(ubody.Age)
		if err != nil {
			c.Writer.WriteString("Age not int type.")
			return
		}
		if err := dbops.AddUserCredential(ubody.Username, ubody.Password, ubody.Email, age); err != nil {
			c.Writer.WriteString(defs.ErrorDBError.Error)
			return
		}

		id := session.GenerateNewSessionId(ubody.Username)
		su := &defs.SignedUp{Success: true, SessionId: id}

		if resp, err := json.Marshal(su); err != nil {
			c.Writer.WriteString(defs.ErrorInternalFaults.Error)
			return
		} else {
			c.Writer.WriteString(string(resp))
		}
	}
}

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, _ := ioutil.ReadAll(c.Request.Body)
		log.Printf("Signin : %s", res)
		ubody := &defs.UserCredential{}
		if err := json.Unmarshal(res, ubody); err != nil {
			log.Printf("%s", err)
			//io.WriteString(w, "wrong")
			c.Writer.WriteString(defs.ErrorRequestBodyParseFailed.Error)
			return
		}

		// Validate the request body

		uname := c.Param("username")
		fmt.Println(len(uname)==0)
		log.Printf("Login url name: %s", uname)
		log.Printf("Login body name: %s", ubody.Username)
		if uname != ubody.Username {
			c.Writer.WriteString(defs.ErrorNotAuthUser.Error)
			return
		}

		log.Printf("%s", ubody.Username)
		pwd, _, _, err := dbops.GetUserCredential(ubody.Username)
		log.Printf("Login pwd: %s", pwd)
		log.Printf("Login body pwd: %s", ubody.Password)
		if err != nil || len(pwd) == 0 || pwd != ubody.Password {
			c.Writer.WriteString(defs.ErrorNotAuthUser.Error)
			return
		}

		id := session.GenerateNewSessionId(ubody.Username)
		si := &defs.SignedIn{Success: true, SessionId: id}
		if resp, err := json.Marshal(si); err != nil {
			c.Writer.WriteString(defs.ErrorInternalFaults.Error)
		} else {
			c.Writer.WriteString(string(resp))
		}

	}
}


func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateUser(c.Writer, c.Request) {
			log.Printf("Unathorized user \n")
			return
		}

		uname := c.Params.ByName("username")
		fmt.Println("66666666666666", uname)
		u, err := dbops.GetUser(uname)
		if err != nil {
			log.Printf("Error in GetUserInfo: %s", err)
			return
		}

		ui := &defs.UserInfo{Id: u.Id}
		if resp, err := json.Marshal(ui); err != nil {
			c.Writer.WriteString(defs.ErrorInternalFaults.Error)
			return
		}else {
			c.Writer.WriteString(string(resp))
		}
	}
}


func AddNewVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateUser(c.Writer, c.Request) {
			log.Printf("Unathorized user \n")
			return
		}

		res, _ := ioutil.ReadAll(c.Request.Body)
		nvbody := &defs.NewVideo{}
		if err := json.Unmarshal(res, nvbody); err != nil {
			log.Printf("Errpr om AddNewVideo: %s", err)
			c.Writer.WriteString(defs.ErrorRequestBodyParseFailed.Error)
			return
		}

		vi, err := dbops.AddVideoInfo(nvbody.AuthorId, nvbody.Name)
		log.Printf("Author id : %d, name : %s", nvbody.AuthorId, nvbody.Name)
		if err != nil {
			log.Printf("Error in AddNewVideo: %s", err)
			c.Writer.WriteString(defs.ErrorDBError.Error)
			return
		}

		if resp, err := json.Marshal(vi); err != nil {
			c.Writer.WriteString(defs.ErrorInternalFaults.Error)
			return
		}else {
			c.Writer.WriteString(string(resp))
		}
	}
}

func ListAllVideos() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateUser(c.Writer, c.Request) {
			return
		}

		uname := c.Params.ByName("username")
		vs, err := dbops.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
		if err != nil {
			log.Printf("Error in ListAllvideos: %s", err)
			c.Writer.WriteString(defs.ErrorDBError.Error)
			return
		}

		vsi := &defs.VideosInfo{vs}
		if resp, err := json.Marshal(vsi); err != nil {
			c.Writer.WriteString(defs.ErrorInternalFaults.Error)
			return
		}else {
			c.Writer.WriteString(string(resp))
		}
	}
}

func DeleteVideo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateUser(c.Writer, c.Request) {
			log.Printf("delete video (((((((((")
			return
		}

		vid := c.Params.ByName("vid")
		err := dbops.DelVideoInfo(vid)
		if err != nil {
			log.Printf("Error in DeletVideo: %s", err)
			c.Writer.WriteString(defs.ErrorDBError.Error)
			return
		}

		go utils.SendDeleteVideoRequest(vid)

		c.Writer.WriteString("Finish Del Video.")
	}
}

func PostComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateUser(c.Writer, c.Request) {
			return
		}

		reqBody, _ := ioutil.ReadAll(c.Request.Body)

		cbody := &defs.NewComment{}
		if err := json.Unmarshal(reqBody, cbody); err != nil {
			c.Writer.WriteString(defs.ErrorRequestBodyParseFailed.Error)
			return
		}

		vid := c.Params.ByName("vid-id")
		if err := dbops.AddNewComments(vid, cbody.AuthorId, cbody.Content); err != nil {
			log.Printf("Error in PostComment: %s", err)
			c.Writer.WriteString(defs.ErrorDBError.Error)
			return
		} else {
			c.Writer.WriteString("ok")
		}
	}
}

func ShowComments() gin.HandlerFunc{
	return func(c *gin.Context) {
		if !ValidateUser(c.Writer, c.Request) {
			return
		}

		vid := c.Params.ByName("vid-id")
		cm, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestampSec())
		if err != nil {
			log.Printf("Error in ShowComments: %s", err)
			c.Writer.WriteString(defs.ErrorDBError.Error)
			return
		}

		cms := &defs.Comments{Comments: cm}
		if resp, err := json.Marshal(cms); err != nil {
			c.Writer.WriteString(defs.ErrorInternalFaults.Error)
		} else {
			c.Writer.WriteString(string(resp))
		}
	}
}
