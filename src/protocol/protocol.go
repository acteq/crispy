package protocol

import (
	// "errors"
)

var msgHanlders map[byte]func([]byte) []byte = make(map[byte]func([]byte) []byte)

func init() {
	
}

func ProcessingData(data []byte) (resp []byte) {
		
	handle, ok := msgHanlders[data[0]]
	if ok == true {
        resp = handle(data[1:])
    } else {
        
	}
	
	return 
}
