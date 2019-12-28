package protocol

import (
	"log"
)

const HeadFlag0 byte = 0xA5
const HeadFlag1 byte = 0x5A
const TailFlag0 byte = 0xCC   
const TailFlag1 byte = 0x33
const TailFlag2 byte = 0xC3
const TailFlag3 byte = 0x3C

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
func  Unpack(data []byte) (consumed int, messages [][]byte ) {
	const HEAD_LEN = 2
	const TAIL_LEN = 4
	var heads = []int{}
	var tails = []int{}
	
	var i int
	for i=0; i < len(data) - TAIL_LEN + 1;  {
		if ( data[i] == HeadFlag0 && data[i+1] == HeadFlag1) {
			heads = append(heads, i)
			i += HEAD_LEN
			consumed = i 
			continue
		}
		if ( data[i] == TailFlag0 && data[i+1] == TailFlag1 && data[i+2] == TailFlag2 && data[i+3] == TailFlag3 ) {
			tails = append(tails, i)
			i += TAIL_LEN
			consumed = i
			continue
		}
		i++
	}
	// log.Printf("unpack %d bytes, loop i=%d, cosumed %d\n", len(data), i, consumed)
// a55a0c0590fae201130c1c0c1530110000000008950000000005dc0000000065191621cc33c33c
	headsIndex := 0
	headsLen := len(heads)
	for _, pos := range tails {
		pairedHead := -1
		for ; headsIndex < headsLen && heads[headsIndex] < pos ; headsIndex ++ {
			pairedHead = heads[headsIndex]
		}
		if pairedHead >=0 {
			checksum := data[pos-2]
			if checksum == countChecksum(data[pairedHead: pos -2]) {
				message := data[pairedHead + HEAD_LEN : pos -2] 
				messages = append(messages, message)
            }else{
                log.Printf("checksum(%x) error for data frame:%s\n", checksum, BytetoH(data[pairedHead: pos + TAIL_LEN] ))
			}
		}
	}
	
	// data frame is broken or uncompleted, left one byte
	if consumed == 0 {
		consumed = len(data) - HEAD_LEN + 1
	}
	return consumed, messages
}

func Pack(message []byte) []byte {
	var bodyLength = byte(len(message))
	var result = make([]byte, 2, bodyLength + 2 + 1 + 1 + 4 )
	//帧头 a55a
	result[0] = HeadFlag0
	result[1] = HeadFlag1

	result = append(result, message...)

	//计算校验和
	var checksum = countChecksum(result)
	//计算长度，值从帧头开始，到校验和前(不算校验和)的值。计算格式：信息体前部长度+信息体
	//长度是从1开始计算
	//校验和, 长度, 帧尾 
	result = append(result, checksum, bodyLength + 2, 0xcc, 0x33, 0xc3, 0x3c)

	return result

}

func CompareChecksum(frame []byte) bool {
	length := len(frame)
	//比对校验和
	var checksum = frame[length-6]

	return  checksum == countChecksum(frame[:length-6])
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
