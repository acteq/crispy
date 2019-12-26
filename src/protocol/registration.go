package protocol

import (
	// "errors"
	"time"
	"strconv"
)

const (
	MsgRegistration     byte = 0x60
	MsgRegistrationRspd byte = 0x61
)

func init() {
	value, ok := msgHanlders[MsgRegistration]
	if ok == true {
        
    } else {
        msgHanlders[MsgRegistration] = registrationHandle
    }
}

func registrationHandle(msg []byte) []byte {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]
	// var  manufacturerCode = msg[5]
	register(towerCraneId, deviceId)

	var resp = make([]byte, 2, 2 + len(deviceId) + 8 )
	// resp type
	resp[0] = MsgRegistrationRspd
	// crane id 
	resp[1] = towerCraneId;

	// sensor device id
	resp = append(resp, deviceId...)

	//registration result: 1 success , 0 failed
	resp = append(resp, 1)

	now := time.Now()
	year, month, day := now.Date() 
	hour, min, sec := now.Clock()

	// time format: yyMMddHHmmss
	fullYear := []byte(strconv.Itoa(year))
	shortYear := fullYear[len(fullYear)-2:len(fullYear)]
	resp = append(resp, shortYear...)
	resp = append(resp, byte(month), byte(day), byte(hour), byte(min), byte(sec))

	// sensor upload realtime data interval, default 15s 
	resp = append(resp, 15)
	// if (black_box_id.equals("9504330")||black_box_id.equals("9504331")||black_box_id.equals("9504333")) {message_segment[7] = 5;}//少量特殊需求设置为5秒一次。
	register(towerCraneId, deviceId)
	return resp

}

/*
 resgister resp
*/
func register(towerCraneId byte,  deviceId []byte) {
	
}
