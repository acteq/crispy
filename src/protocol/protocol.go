package protocol

import (
	// "errors"
	"time"
	"log"
)

var msgHanlders map[byte]func([]byte) []byte = make(map[byte]func([]byte) []byte)

func init() {
	
}

func HandleMessage(msg []byte) (resp []byte) {
	if len(msg) < 1 {
		return nil
	}

	handle, ok := msgHanlders[msg[0]]
	if ok == true {
        resp = handle(msg)
    } else {
        log.Printf("unknown frame %2X\n", msg[0])
	}
	
	return 
}


func BytesToTime(buf []byte) (result time.Time) {
	if len(buf) < 6 {
		return 
	}
	return time.Date(2000 + int(buf[0]), time.Month(buf[1]), int(buf[2]), int(buf[3]), int(buf[4]), int(buf[5]), 0,  time.UTC)
}