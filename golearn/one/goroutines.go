package main

import (
	"fmt"
	"time"
	
)

//goroutine is a lightweight thread of execution

//suppose we have a nornmal function f through  synchronously.
func f(from string){
	for i:=0;i < 3;i++{
		fmt.Println(from,":", i)
	}
}

func Goroutines(){
	f("yollow")
	go f("manly")
	go func(msg string) {
        fmt.Println(msg)
    }("going")
	time.Sleep(time.Second)
    fmt.Println("done")
}