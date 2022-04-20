package defs

//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"pwd"`
	Email string `json:"email"`
	Age string `json:"age"`
}

//Data model

type User struct {
	Id int
	LoginName string
	Pwd string
}

type UserInfo struct {
	Id int `json:"id"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

//COmment model
type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

// Session model
type SimpleSession struct {
	Username string
	TTL int64 	// Time To Live
}

//response
type SignedIn struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`

}