package main

import (
	"fmt"
	"time"
)

//Go 语言支持并发，通过go关键字开启 goroutine 即可
//goroutine是轻量级线程（学名：协程），goroutine的调度是由 Golang 运行时进行管理
//goroutine 语法格式： go 函数名( 参数列表 )

func main() {
	// 输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行
	go say("world")
	say("hello")

	// 测试通道
	s := []int{7, 2, 8, -9, 4, 0}
	// 无缓冲channel：发送和接收需要同步，发送阻塞直到数据被接收，接收阻塞直到读到数据
	c := make(chan int)
	go sum(s[:3], c)
	go sum(s[3:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)

	bufferChannel()

	// 有缓冲channel：不要求发送和接收操作同步，当缓冲满时发送阻塞，当缓冲空时接收阻塞
	c2 := make(chan int, 10)
	go listAndClose(cap(c2), c2)
	// range 函数遍历每个从通道接收到的数据
	// 因为c2在发送完10个数据之后就关闭了通道，所以range函数在接收到10个数据之后就结束。
	// 如果上面的c2通道不关闭，那么range函数就不会结束，从而在接收第 11 个数据的时候就阻塞
	for i := range c2 {
		fmt.Println(i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 通道: chan关键字，通道在使用前必须先创建：ch := make(chan int)
// ch <- v    // 把 v 发送到通道 ch
// v := <-ch  // 从 ch 接收数据,并把值赋给 v
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

// 带缓冲区通道
func bufferChannel() {
	// 缓冲区大小为2
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据，而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 遍历和关闭通道
func listAndClose(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
