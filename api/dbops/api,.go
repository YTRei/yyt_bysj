package dbops

import (
	"bysj_VEDIO/api/defs"
	"bysj_VEDIO/api/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)


// Handle of User

func AddUserCredential(loginName string, pwd string, email string, age int) error  {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd, email, age) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Error 1 : ", err)
		return err
	}
	isSave, _ := UserHadSave(loginName)
	if !isSave{
		fmt.Println("Error 2 : ")
		return nil
	}
	_, err = stmtIns.Exec(loginName, pwd, email, age)
	if err != nil {
		fmt.Println("Error 3 : ", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

func UserHadSave(loginName string) (bool, error) {
	pwd, email, _, err := GetUserCredential(loginName)
	if err != nil {
		log.Printf("Error of ReGetUser : %v", err)
		return false, err
	}
	if pwd != "" || email != ""  {
		log.Printf("Deleting user test failed")
		return false, err
	}
	return true, nil
}

func GetUserCredential(loginName string) (string, string, string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd, email, age FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", "", "", err
	}

	var pwd string
	var email string
	var age string
	err = stmtOut.QueryRow(loginName).Scan(&pwd, &email, &age)
	if err != nil && err != sql.ErrNoRows {	//没有结果
		return "", "", "", err
	}
	defer stmtOut.Close()

	return pwd, email, age, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUSer error : %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		fmt.Println(111111111111111)
		return nil, err
	}

	var id int
	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}

	defer stmtOut.Close()

	return res, nil
}

//Hand with video

func AddVideoInfo(aid int, name string) (*defs.VideoInfo, error) {
	//create UUID
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
    (id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string)(*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var name string
	var dct string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: dct,
	}

	return res, nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		INNER JOIN users ON video_info.author_id = users.id
		WHERE users.login_name = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video_info.create_time DESC`)


	var res []*defs.VideoInfo

	if err != nil {
		return res, err
	}

	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	defer stmtOut.Close()

	return res, nil
}

func DelVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		log.Printf("DeleteVideo error : %s", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}


//comment handle

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, comment) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments
    INNER JOIN users ON comments.author_id = users.id
	WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
	ORDER BY comments.time DESC`)
	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{
			Id:       id,
			VideoId:  vid,
			Author: name,
			Content:  content,
		}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}