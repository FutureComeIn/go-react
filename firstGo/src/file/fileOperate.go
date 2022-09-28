package main

import (
	"fmt"
	"os"
)

func main() {
	// directory()
	// file()
	readFile()
}
/*
打开文件
	1、func Open(name string) (file *File, err Error)
	   该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
	2、func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
	   打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
*/
func file() {
	userFile := "test/1.txt"
	// 根据文件名创建文件，返回一个文件对象，默认权限是0666，可读写
	out, err := os.Create(userFile)
	// 根据文件描述符创建相应的文件，返回一个文件对象
	os.NewFile(1, "test/2.txt")
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer out.Close()

	for i := 0; i < 10; i++ {
		out.WriteString("Just a test!\r\n")
		out.Write([]byte("Just a test!\r\n"))
	}
}

/**
	func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
	从off开始读取数据到b中
 */
func readFile(){
	userFile := "test/1.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()

	buf := make([]byte, 1024)
	for {
		// 读取数据到buf中
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func directory() {
	// 创建名称为name的目录，权限设置是perm，例如0777
	os.Mkdir("test", 0777)
	os.MkdirAll("test/test1/test2", 0777)

	// 删除没子级的目录或文件
	err := os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("test")
}
