package main

import (
	"fmt"
	"net"
	"os"
)

// os.Args：带参数运行go run netProgram.go 123.23.23，则os.Args[1]=123.23.23，默认有os.Args[0]=该程序路径
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
