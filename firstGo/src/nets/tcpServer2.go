package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// 和tcpServer不同之处：建立了长连接，可以一直读取客户端内容
func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// 设置超时，当一定时间内客户端无请求发送，conn便会自动关闭
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// set maxium request length to 128B to prevent flood attack
	request := make([]byte, 128)
	// close connection before exit
	defer conn.Close()
	for {
		// 不断读取客户端发来的请求
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			// connection already closed by client
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		// clear last read content，因为conn.Read()会将新读取到的内容append到原内容之后
		request = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
