package main

import (
	"fmt"
	"runtime"
	"strings"
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
	//timestamp := time.Now().UnixNano() / 1e6
	//appKey := "8aef559a8dbe6000"
	//md5Str := "appKey=" + appKey + "&timestamp=" + strconv.Itoa(int(timestamp)) + "&random=" + "222"
	//h := md5.New()
	//h.Write([]byte(md5Str))
	//md5 := hex.EncodeToString(h.Sum(nil))
	//
	//fmt.Println(timestamp)
	//fmt.Println(md5)

	fmt.Println(strings.TrimRight("1,2,3,", ","))
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
