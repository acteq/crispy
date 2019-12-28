package main

import (
    "fmt"
	"log"
	"os"
    "net"
	"time"
	"tower-crane/protocol"
)

// sockopt设置接口
// SetKeepAlive 是否开启长连接
// SetKeepAlivePeriod 设置长连接的周期，超出会断开
// SetLinger 设定当连接中仍有数据等待发送或接受时的Close方法的行为。
// SetNoDelay （默认no delay） 设定操作系统是否应该延迟数据包传递，以便发送更少的数据包（Nagle's算法）。默认为真，即数据应该在Write方法后立刻发送。
// SetWriteBuffer 连接的系统发送缓冲
// SetReadBuffer 连接的系统接收缓冲
// 标准TCP层协议里把对方超时设为2小时，若服务器端超过了2小时还没收到客户的信息，它就发送探测报文段，若发送了10个探测报文段（每一个相隔75S）还没有收到响应，就假定客户出了故障，并终止这个连接。因此应对tcp长连接进行保活。

func init()  {
    log.SetOutput(os.Stdout)
    log.SetFlags(log.Llongfile)	
}

const PORT = 21579

func main() {
    listen, err := net.Listen("tcp", ":21579")
    if err != nil {
		fmt.Println("listen error:", err)
		log.Println("error listen:", err)
        return
    }

	defer listen.Close()
    log.Println("listen ok")
    for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            break
        }
        // start a new goroutine to handle
        // the new connection.
        go handleConn(conn)
    }
}

func handleConn(c net.Conn) {
	defer c.Close()

	tcpConn, ok := c.(*net.TCPConn)
	if !ok {
		//error handle
	}

	tcpConn.SetNoDelay(true)
    
	// read from the connection
    var bufRead = make([]byte, 65536)
    var left = 0
    for {
        // var bufRead = buf
        c.SetReadDeadline(time.Now().Add(time.Second * 300))
        n, err := c.Read(bufRead[left:])
        if err != nil {
            log.Printf("conn read %d bytes,  error: %s", n, err)
            if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
                continue
            }
            return
        }
        // log.Printf("read %d bytes, content is %s\n", n, protocol.BytetoH(bufRead[:n]))
        consumed, messages := protocol.Unpack(bufRead[:left + n])
        if left + n - consumed > 0 {
            log.Printf("last %d bytes, received %d bytes, consumed %d bytes, left %d bytes, messages %d\n", left, n, consumed, left + n - consumed, len(messages))
            log.Printf("last data %s\n", protocol.BytetoH(bufRead[:left]))
            log.Printf("received data:%s\n", protocol.BytetoH(bufRead[left: left + n]))
            log.Printf("left data:%s\n", protocol.BytetoH(bufRead[consumed:left + n]))
            copy(bufRead, bufRead[consumed:left + n])
        }
        left = left + n - consumed
        for _, message := range messages {
            resp := protocol.HandleMessage(message)
            if resp == nil {
                continue
            }
            packed := protocol.Pack(resp)
            n, err = c.Write(packed)
            if err != nil {
                log.Println("conn write error:", err)
            } else {
                log.Printf("write %d bytes: %s\n", n, protocol.BytetoH(packed[:n]))
            }
        }
    }
}


