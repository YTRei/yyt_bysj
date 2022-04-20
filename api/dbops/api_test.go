package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate vadeo_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDelUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("Rei", "123", "66@163.com", 15)
	if err != nil {
		t.Errorf("Error of AddUser : %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, email, age, err := GetUserCredential("Rei")
	if pwd != "123" || err != nil || email != "66@163.com" || age != 15{
		t.Errorf("Error of GetUser : %v", err)
	}

}

func testDelUser(t *testing.T) {
	err := DeleteUser("Rei", "123")
	if err != nil {
		t.Errorf("Error of DelUser : %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, email, age, err := GetUserCredential("Rei")
	if err != nil {
		t.Errorf("Error of ReGetUser : %v", err)
	}

	if pwd != "" || email != "" || age == 15{
		t.Errorf("Deleting user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddVideoInfo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DelVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments",testAddNewComments)
	t.Run("ListComments", TestListComments)
}

func testAddNewComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "hey!"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddNewComments: %v", err)
	}
}

func TestListComments(t *testing.T) {
	vid := "12345"
	from := 151476480
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments : %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}