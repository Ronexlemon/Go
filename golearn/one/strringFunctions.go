package main

import (
	"fmt"
	ss "strings"
)

var p = fmt.Println

func StringFunctions(){
	p("Contains",ss.Contains("test","st"))
	p("Count",ss.Count("test","t"))
	p("HasPrefix",ss.HasPrefix("test","te"))
	p("HasSuffix",ss.HasSuffix("test","st"))
	p("Index",ss.Index("test","s"))
	p("Joins",ss.Join([]string{"test","it"},"-"))
	p("Repeat",ss.Repeat("test",3))
	p("Replace",ss.Replace("test","t","u",-1))
	p("Replace",ss.Replace("test","t","u",1))
	p("Split",ss.Split("t-e-s-t","-"))
	p("ToLower",ss.ToLower("YOLLOW"))
	p("ToUpper",ss.ToUpper("test"))
}