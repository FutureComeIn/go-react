package main

import (
	"html/template"
	"os"
)

func main() {
	template2()
}

type People struct {
	UserName string
	email	string  //首字母是小写，则不会导出
}
func template1(){
	t := template.New("hello.html")
	t, _ = t.Parse("hello {{.UserName}}!{{.email}}")
	p := People{UserName: "Tom",email: "42424"}
	t.Execute(os.Stdout, p)
}

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}
func template2(){
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}