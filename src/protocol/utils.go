package protocol

import (
    "bytes"
    "io/ioutil"

    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"

    "fmt"
    "strconv"
)

func HextoB(str string) []byte { 
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


func Utf82Gb2312(s []byte)([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}