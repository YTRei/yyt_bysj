package dbops

import (
	"bysj_VEDIO/api/defs"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func InserSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtInts, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	fmt.Println("input: \n",sid, "\n",ttlstr,"\n", uname)
	if err != nil {
		return err
	}

	_, err = stmtInts.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	fmt.Println("Add session into DB")
	defer stmtInts.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	if err != nil {
		return nil, err
	}

	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("444444444444444")
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	}else{
		fmt.Println("5555555555", err)
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string

		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("%s", err)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{
				Username: login_name,
				TTL:      ttl,
			}
			m.Store(id, ss)
			log.Printf("session id : %s , ttl : %v ", id, ss.TTL)
		}

	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOUt, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err := stmtOUt.Query(sid); err != nil {
		return err
	}

	return nil
}