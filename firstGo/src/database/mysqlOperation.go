package main

import (
	"database/sql"
	"fmt"
	// MySQL驱动： https://github.com/go-sql-driver/mysql 支持database/sql，全部采用go写
	// 驱动有问题：下载这个https://github.com/go-sql-driver/mysql/releases/tag/v1.4.0
	// _的意思是引入后面的包名而不直接使用这个包中定义的函数，变量等资源（包在引入的时候会自动调用包的init函数以完成对包的初始化）
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 可以不指定地址创建
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("Jack", "财务部门", "2022-05-17")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("Tom2", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
