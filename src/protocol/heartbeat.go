package protocol

import (
	// "errors"
	"log"
)

const (
	MsgHeartBeat    byte = 0x80 //心跳包
	MsgHeartBeatRspd    byte = 0x8e
)
func init() {
	_, ok := msgHanlders[MsgHeartBeat]
	if ok == true {
		log.Printf("can't add function heartbeatHandle, the key %2X has exits\n", MsgHeartBeat)   
    } else {
        msgHanlders[MsgHeartBeat] = heartbeatHandle
    }
}

func heartbeatHandle(msg []byte) []byte {
	if len(msg) < 6 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}

	//塔吊编号
	towerCraneId := msg[1]
	//设备编号 
	deviceId := msg[2:5]
	
	interval := msg[5]

	return heartBeatRspd(towerCraneId, deviceId, interval)
}

func heartBeatRspd(towerCraneId byte,  deviceId []byte, interval byte) []byte {
	log.Printf("heartbeat\n")
	var message = make([]byte, 1 + 1 + len(deviceId) + 1 )
	// message type
	message[0] = MsgHeartBeatRspd

	// crane id
	message[1] = towerCraneId

	// sensor device id
	for i, val := range  deviceId {
		message[2+i] = val
	}
	
	message[len(message) - 1] = interval

	return message
}
