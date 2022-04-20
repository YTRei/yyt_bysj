package taskrunner

import (
	"bysj_VEDIO/scheduler/dbops"
	"errors"
	"fmt"
	"log"
	//"os"
	"sync"
	"bysj_VEDIO/scheduler/ossops"
)

func deleteVideo(vid string) error {
	//tmp, _ := os.Getwd()

	//pp := tmp + VIDEO_PATH + vid
	//fmt.Println("now: " + pp)

	// 本地
	//err := os.Remove(VIDEO_PATH + vid)
	//
	//if err != nil && !os.IsNotExist(err){
	//	log.Printf("Deleting video error: %v", err)
	//	return err
	//}
	//
	//return nil

	ossfn := "video/" + vid
	bn := "yuan-videos"
	ok := ossops.DeleteObject(ossfn, bn)

	if !ok {
		log.Printf("Deleting video error, oss operation failed")
		return errors.New("Deleting video error.")
	}

	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	//fmt.Println("qqq")
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		//fmt.Println("err 1 ")
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		//fmt.Println("err 2 ")
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
				case vid :=<- dc:
					go func(id interface{}) {
						if err := deleteVideo(id.(string)); err != nil {
							fmt.Println("panic 1")
							errMap.Store(id, err)
							return
						}
						if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
							fmt.Println("panic 2")
							errMap.Store(id, err)
							return
						}
					}(vid)
			default:
				break forloop
			}
		}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			fmt.Println("panic 3")
			return false
		}
		return true
	})

	return err
}