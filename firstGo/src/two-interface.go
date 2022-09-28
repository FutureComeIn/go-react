package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

type IPhone struct {
}

// 接口实现类定义
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main(){
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}