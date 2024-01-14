package main

import (
	"fmt"
	"time"
)

func Time(){
	p:=fmt.Println

	now:= time.Now()
	p(now)

	then:= time.Date(2010,11,20,10,56,35,4343533535353,time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
}

