package protocol

import (
	"testing"
	"bytes"
	"strings"
	"fmt"
	"strconv"
)

var messages = map[string]string{
	"60 02 90 F7 2F 01":"A5 5A 60 02 90 F7 2F 01 18 08 CC 33 C3 3C",  //Registration
	"0c 04 BC 61 4E 01 13 07 18 11 19 06 11 00 00 00 00 0b cf 03 cf 00 00 05 dc 10 00 00 00 65 09": "a5 5a 0c 04 BC614E 01 130718111906 11 0000 0000 0b cf 03 cf 00 00 05 dc 10 00 00 00 65 09 9f 21 cc 33 c3 3c", //RealtimeData
	"3D 02 90 F7 2F 02 11 13 01 08 0A 1C 21 13 01 08 0A 1D 21 17 D3 00 C7 01 2C 00 82 00 00 00 00 00 00 00 00 00 82 00 00 00 00 01 04" : "A5 5A 3D 02 90 F7 2F 02 11 13 01 08 0A 1C 21 13 01 08 0A 1D 21 17 D3 00 C7 01 2C 00 82 00 00 00 00 00 00 00 00 00 82 00 00 00 00 01 04 B5 2D CC 33 C3 3C", //CycleInfo
	"3d 01 90 fb f1 01 11 13 07 12 10 0b 15 13 07 12 10 10 2d 04 85 00 eb 00 c7 00 5b 00 fd 00 54 05 69 00 fc 00 a3 08 87 00 b3 00 5b" : "a5 5a 3d 01 90 fb f1 01 11 13 07 12 10 0b 15 13 07 12 10 10 2d 04 85 00 eb 00 c7 00 5b 00 fd 00 54 05 69 00 fc 00 a3 08 87 00 b3 00 5b 31 2d cc 33 c3 3c", //CycleInfo2
	"80 02 90 F7 2F 3C" : "A5 5A 80 02 90 F7 2F 3C 73 08 CC 33 C3 3C", //HeartBeat
	"8E 02 90 F7 2F 3C" : "A5 5A 8E 02 90 F7 2F 3C 81 08 CC 33 C3 3C", //HeartBeatResponse
	
}


type TestCase struct {
	message []byte
	frame []byte
}

func Hextob(str string) []byte { 
	slen:=len(str) 
	bHex:=make([]byte,slen/2) 
	ii :=0 
	for i:=0;i<len(str);i=i+2 { 
		if slen!=1 { 
			ss := string(str[i]) + string(str[i+1])
			bt,_ := strconv.ParseInt(ss,16,32) 
			bHex[ii] = byte(bt)
			ii = ii+1
			slen=slen-2
		} 
	}
	return bHex;
} 
	/*字节数组转16进制可以直接使用 fmt自带的*/ 
func BytetoH(b []byte) string { 
	return fmt.Sprintf("%x",b)
}
	
func convert2TestCases(msgMap map[string]string) []TestCase {

	var testcases = make([]TestCase, 0, len(msgMap))

	for msg, frame := range msgMap {

		dst := Hextob(strings.ReplaceAll(msg, " ", ""))
		// fmt.Printf("%s\n", BytetoH(dst))
		// fmt.Printf("%s\n",strings.ReplaceAll(msg, " ", ""))
		dst2 := Hextob(strings.ReplaceAll(frame, " ", ""))
		// fmt.Printf("%s\n", BytetoH(dst2))
		// fmt.Printf("%s\n",strings.ReplaceAll(frame, " ", ""))
		testcases = append(testcases, TestCase{dst, dst2})
	}
	return testcases
}


func TestCountChecksum(t *testing.T){
	var testcases = convert2TestCases(messages)
    for _, tc := range testcases {
		result := countChecksum(tc.frame[:len(tc.frame)-6])
        if  result != tc.frame[len(tc.frame)-6] {
			t.Errorf("countChecksum failed: %s -- %x -- %x", BytetoH(tc.frame), tc.frame[len(tc.frame)-6] ,result )
        }
	}
	
}

func TestPack(t *testing.T) {
	var testcases = convert2TestCases(messages)
    for _, tc := range testcases {
		result := Pack(tc.message)
        if ! bytes.Equal(result, tc.frame) {
			t.Errorf("pack failed: %s -- %s" , BytetoH(result) ,BytetoH(tc.frame)   )
        }
	}
}

func TestUnpack(t *testing.T) {
	var testcases = convert2TestCases(messages)
    for _, tc := range testcases {
		_, result, err := Unpack(tc.frame)
		if err != nil {
			t.Errorf("unpack failed: %s", err )
			continue
		}
        if ! bytes.Equal(result, tc.message) {
			t.Errorf("unpack failed" )
		}
	}
}


