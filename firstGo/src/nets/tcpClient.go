package main

import (
	"fmt"
	"net"
	"os"
)

// 模拟一个基于HTTP协议的客户端请求去连接一个Web服务端
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	// 获取一个TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	// 建立一个TCP连接，并返回一个TCPConn类型的对象
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	// 写http请求头，发送请求信息
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	// 读取
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)
	// 从conn中读取全部的文本，也就是服务端响应反馈的信息
	// result, err := ioutil.ReadAll(conn)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
