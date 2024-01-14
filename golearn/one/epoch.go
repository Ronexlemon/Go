package main

import (
	"fmt"
	"time"
)
var pp = fmt.Println

func Epoch(){
	now:= time.Now()
	pp(now)

	pp(now.Unix())
	pp(now.UnixMilli())
	pp(now.UnixNano())

	pp(time.Unix(now.Unix(),0))
	pp(time.Unix(0,now.UnixNano()))

}