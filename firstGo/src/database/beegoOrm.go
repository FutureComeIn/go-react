package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// beego orm：用Go进行ORM操作的库，实现了struct到数据表记录的映射，是一个十分轻量级的Go ORM框架，
// 即操作数据库不直接在代码中用原生sql

// 定义的struct添加到orm model后，会帮你创建为表，需驼峰命名（会自动转化成下划线字段，如定义struct名字UserInfo，转化成表是user_info，字段命名也遵循）
type User struct {
	Id      int
	Name    string   `orm:"size(60)"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation，User带外键字段profile_id
	// Post    []*Post `orm:"reverse(many)"`  // 还能一对多，设置一对多的反向关系
}
type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
}

type Profile struct {
	Pid  int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键，但要自己改成 AUTO_INCREMENT
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/test?charset=utf8", 30)
	// 根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	// 根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)

	// 注册定义的 model，根据struct创建表
	//orm.RegisterModel(new(User))
	// 也可以同时注册多个 model
	orm.RegisterModel(new(User), new(Profile))

	// 创建 table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	profileOperate(o)
	//userOperate(o)
}

func profileOperate(o orm.Ormer) {
	//profile := Profile{Age: 24}
	// id, err := o.Insert(&profile)
	// fmt.Println("profileId：", id, "，err：", err)
	profile := Profile{Pid: 1}
	err := o.Read(&profile)
	fmt.Println("profile:{}",profile,",error:{}",err)

}

func userOperate(o orm.Ormer) {
	profile := Profile{Pid: 2}
	user := User{Name: "Tom", Profile: &profile}
	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 同时插入多个对象:InsertMulti：insert into table (name, age) values("slene", 28),("astaxie", 30),("unknown", 20)
	users := []User{
		{Name: "张三", Profile: nil},
		{Name: "李四", Profile: nil},
		{Name: "王五", Profile: nil},
	}
	// 第一个参数 bulk 为并列插入的数量，第二个为对象的slice，bulk 为 1 时，将会顺序插入 slice 中的数据
	successNums, err := o.InsertMulti(1, users)
	fmt.Println("成功插入的数量：", successNums)

	// 更新表
	//user.Name = "Jack2"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
