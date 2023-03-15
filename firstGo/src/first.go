package main

import "fmt"

func main() {
	// 变量定义
	var a = "hello"
	b := 4
	var c bool = true

	// 常量定义
	const LENGTH int = 10

	// 多重赋值
	const aa, bb, cc = 1, false, "str"

	// iota：起const索引作用
	const (
		a1 = iota //0
		b1        //1
		c1        //2
		d1 = "ha" //独立值，iota += 1
		e1        //"ha"   iota += 1
		f1 = 100  //iota +=1
		g1        //100  iota +=1
		h1 = iota //7,恢复计数
		i1        //8
	)

	fmt.Print("first " + a + "\n")
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1)

	if b < 20 {
		fmt.Printf("b 小于 20\n")
	} else if b < 10 {
		fmt.Println("b 小于 10")
	} else {
		fmt.Println("b 大于 10")
	}

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 无while
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// For-each range 循环
	strs := []string{"google", "runoob"}
	for i, s := range strs {
		fmt.Println(i, s)
	}

	x, y := swap("hello", "Tom")
	fmt.Println(x, y)
	pointer()

	// 忽略的字段为 0 或 空
	book1 := Book{title: "Go 语言", author: "www.runoob.com"}
	book2 := Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book2.author)

	mapUse()
	typeTransfer()

}

// 函数定义：func function_name( [parameter list] ) [return_types] {}
func swap(x, y string) (string, string) {
	return y, x
}

// 数组和切片
func arrAndSlice() {
	// 数组定义 加三点...可能是会根据值自动确定长度吧
	arr1 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	arr2 := [3]float32{1000.0, 2.0, 3.4}
	// 输出每个数组元素的值
	for k := 0; k < len(arr2); k++ {
		fmt.Printf("arr1[%d] = %f,arr2[%d] = %f\n", k, arr1[k], k, arr2[k])
	}
	fmt.Println("数组：", arr2[1:3])

	// 切片：动态数组，建的时候长度空着，append追加，copy
	num1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	num2 := make([]int, 0, 5) // make([]type, length, capacity)，capacity可选参数
	num2 = append(num2, 2)
	// 拷贝 num2 的内容到 num3（但新切片得是之前切片的两倍容量）
	num3 := make([]int, len(num2), (cap(num2))*2)
	copy(num3, num2)
	fmt.Println("切片1：", num2, num1[2:4], num1[2:], num1[:4])
	fmt.Println("切片2：", num3, num1)
}

// map
func mapUse() {
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India "] = "新德里"
	map2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	for country := range map1 {
		fmt.Println(country, "首都是", map1[country])
	}
	delete(map1, "France")
	for country := range map1 {
		fmt.Println(country, "首都是", map1[country])
	}
	// 查看元素在集合中是否存在
	capital, ok := map1["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	for country := range map2 {
		fmt.Println(country, "首都是", map2[country])
	}
}

// 指针
func pointer() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的存储地址 */

	var ptr *int // 不赋值就是空指针，值为 nil，指代零或空值
	if ptr == nil {
		fmt.Println(ptr) // <nil>
	}

	fmt.Printf("ptr 的值为 : %x\n", ptr)

	fmt.Printf("a 变量的地址是: %x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

func typeTransfer() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}

// 结构体，即类
type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}
