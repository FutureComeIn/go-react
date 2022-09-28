package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/**
 * 说明：
 * 终端运行 go run service.go
 * 浏览器打开：http://localhost:3001/
 */

// service.go 是程序的启动入口，上一章说过里面是一个 Web 服务器
func main() {
	// 1、新建路由器实例
	mux := http.NewServeMux()
	// 2、实例添加一个首页的 Handle Func 函数：接收一个字符串的 pattern 路径表达式和一个 handler 函数
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 输出一行简单的文字
		fmt.Fprintf(writer, "I'm a cook server3.")
	})

	timeOut := time.Second * 45
	srv := &http.Server{
		Addr:           ":3001",
		Handler:        mux,
		ReadTimeout:    timeOut,     // 读 超时时间
		WriteTimeout:   timeOut,     // 写 超时时间
		IdleTimeout:    timeOut * 2, // 空闲 超时时间
		MaxHeaderBytes: 1 << 20,     // 头部数据大小
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf(" listen and serve http server fail:\n %v ", err)
		}
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	<- exit
	ctx, cacel := context.WithTimeout(context.Background(), timeOut)
	// defer是go中一种延迟调用机制，defer跟着的函数直到 return 前才被执行，通常用于释放资源
	defer cacel()

	err := srv.Shutdown(ctx)
	log.Println("shutting down now. ", err)
	os.Exit(0)

	// http.ListenAndServe(":3001", mux)
}
