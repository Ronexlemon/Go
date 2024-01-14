package main

import (
	"fmt"
	"time"
)

func TimeFormating(){
	ff:=fmt.Println
	t:= time.Now()
	ff(t.Format(time.RFC3339))

	t1,e:= time.Parse(time.RFC3339,"2012-11-01T22:08:41+00:00")
	if   e !=nil{
		panic(e)
	}
	p(t1)
	p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2,_:= time.Parse(form, "8 41 PM")
    p(t2)
}