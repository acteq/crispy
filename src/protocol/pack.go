package protocol

import (
	"errors"
	"fmt"
)

const FrameFlag0 byte = 0xA5
const FrameFlag1 byte = 0x5A
const FrameTailFlag0 byte = 0xCC   
const FrameTailFlag1 byte = 0x33
const FrameTailFlag2 byte = 0xC3
const FrameTailFlag3 byte = 0x3C

type DataFrame struct {
	head uint16
	frameType byte
	craneId byte
	deviceId byte
	checksum byte
	length byte
	tail uint32
}
type DataFrameHead struct {
	flag0 int8
	flag1 int8
	frameType byte
	craneId byte
	deviceId byte
}

type DataFrameTail struct {
	checksum byte
	length byte
	tailFlag uint32
}
/**
* 分割数据并交由数据处理类处理
* 设备发送的数据可能会以a55a000000...cc33c33ca55a010000...cc的形式粘连，需要切割。
* @param os
* @param data
*/
func  Unpack(data []byte) (consumed int, message []byte, err error) {
	var length = len(data)
	if length < 13 {
		err = errors.New("size too small")
		return 
	}
	if ( data[0] != FrameFlag0 || data[1] != FrameFlag1) {
		err = errors.New("error head")
		return
	}

	//get craneId, deviceId

	//search the end
	const StartOfSearch = 7
	const LengthOfTail = 4
	var i  = StartOfSearch
	var found = false

	for ; (i + LengthOfTail - 1) <= length  ; i++ {
		if ( data[i] == FrameTailFlag0 && data[i+1] == FrameTailFlag1 && data[i+2] == FrameTailFlag2 && data[i+3] == FrameTailFlag3 ) {
			found = true
			break	
		}
	}

	if !found {
		err = errors.New("not found the tail")
		fmt.Printf("%d, %X%X%X%X%X\n", i, data[i-1], data[i], data[i+1], data[i+2] , data[i+3] )
	}
	
	consumed = i + LengthOfTail + 1
	message = data[ 2 : i -2]

	//比对校验和
	var frame_checksum = data[i-2];

	if( frame_checksum != countChecksum(data[:length-6]) ){
		err = errors.New("checksum error")
	}
	
	return
}


func Pack(message []byte) []byte {
	var bodyLength = byte(len(message))
	var result = make([]byte, 2, bodyLength + 2 + 1 + 1 + 4 )
	//帧头 a55a
	result[0] = FrameFlag0
	result[1] = FrameFlag1

	result = append(result, message...)

	//计算校验和
	var checksum = countChecksum(result)
	//计算长度，值从帧头开始，到校验和前(不算校验和)的值。计算格式：信息体前部长度+信息体
	//长度是从1开始计算
	//校验和, 长度, 帧尾 
	result = append(result, checksum, bodyLength + 2, 0xcc, 0x33, 0xc3, 0x3c)

	return result

}


/**
* 计算校验和
* 累加和校验【每字节相加（16进制）取后末两位】
* @param bytes
* @return
*/
func countChecksum(data []byte) byte {
	var result int = 0
	for i :=0 ; i < len(data); i++ {
		result += int(data[i])
	}
	//a5 5a 0c 04 BC614E 01 130718111906 11 0000 0000 0b cf 03 cf 00 00 05 dc 10 00 00 00 65 09
	return byte(result%256)
}
