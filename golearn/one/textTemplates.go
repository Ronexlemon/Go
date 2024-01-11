package main

import (
	
	"os"
	"text/template"
)

func TextTempplates(){
	t1:= template.New("t1")
	t1,err:=t1.Parse("Value is {{.}}\n")
	if err !=nil{
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))
	t1.Execute(os.Stdout,"some text")
	t1.Execute(os.Stdout,5)
	t1.Execute(os.Stdout,[]string{
		"Go","Rust","C++","C#",
	})

	create:= func(name,t string)*template.Template{
		return template.Must(template.New(name).Parse(t))
	}
	t2:= create("t2","Name: {{.Name}}\n")
	t2.Execute(os.Stdout,struct{Name string}{"Jane doe"})
	t2.Execute(os.Stdout,map[string]string{
		"Name":"Micky",
	})

	t3:= create("t3","{{if.}} yes {{else }}no{{end}}\n")

	t3.Execute(os.Stdout,"no empty")

	t4:= create("t4","Range:{{range .}} {{.}} {{end}}\n")
	t4.Execute(os.Stdout,[]string{"GO","RUST","C++","C"})

}