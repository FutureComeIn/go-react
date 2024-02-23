package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         int      `json:Name`
	Publications []string `json:Publication,omitempty`
}

func main() {
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Println(name, s.Tag)
	}

	// %v输出结构体各成员的值；
	// %+v输出结构体各成员的名称和值；
	// %#v输出结构体名称和结构体各成员的名称和值
	auther := Author{
		Name:         1,
		Publications: []string{"a", "b"},
	}
	fmt.Printf("%#v", auther)
}
