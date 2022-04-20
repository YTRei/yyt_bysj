package session

import (
	"bysj_VEDIO/api/dbops"
	"bysj_VEDIO/api/defs"
	"bysj_VEDIO/api/utils"
	"fmt"
	"sync"
	"time"
)



var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64{
	return time.Now().UnixNano()/1000000
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano()/1000000
	ttl := ct + 30 * 60 * 1000//Serverside session valid time : 30 min

	ss := &defs.SimpleSession{
		Username: un,
		TTL:      ttl,
	}
	fmt.Println("add session !")
	sessionMap.Store(id, ss)
	dbops.InserSession(id, ttl, un)

	return id
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}


func IsSessionExpired(sid string) (string, bool) {
	//fmt.Println("KKKKKKKKKKKKKKKKK")
	ss, ok := sessionMap.Load(sid)
	ct := time.Now().UnixNano()/1000000
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			//fmt.Println(ss.(*defs.SimpleSession).TTL,"\n",ct)
			//fmt.Println("111111111111")

			//delete expried session
			deleteExpiredSession(sid)
			return "",true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		//fmt.Println(ss == nil)
		if err != nil || ss == nil {
			//fmt.Println("2222222222")
			return "", true
		}

		if ss.TTL < ct {
			//fmt.Println("3333333333")
			deleteExpiredSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.Username, false
	}
	//fmt.Println("wwwwwwwww")
	return "", true
}