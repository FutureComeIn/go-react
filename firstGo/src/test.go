package main

import (
	"fmt"
	"runtime"
	"time"
)

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func main() {
	//x := 10
	//y := 20
	//fmt.Println("Before swap:", x, y)
	//swapData(&x, &y)
	//fmt.Println("After swap:", x, y)

	result := divide(10, 0)
	fmt.Println("Result:", result)
	//	执行结果
	//	Error: division by zero
	//	Result: 0

	sample := "我爱GO"
	runeSamp := []rune(sample)
	runeSamp[0] = '你'
	fmt.Println(string(runeSamp)) // "你爱GO"
	fmt.Println(len(runeSamp))    // 4
}

// panic：是Go语言中用于处理异常的机制。当程序遇到无法处理的错误时，可以使用panic引发一个异常，中断程序的正常执行。
// recover：用于捕获并处理panic引发的异常，使程序能够继续执行
func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func swapData(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// 计算日期相差多少月
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}

func test3() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func test2() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}

}

func test1() {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver

	panic("未付")
}
