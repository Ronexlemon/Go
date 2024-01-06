package main

import (
	"fmt"
	"time"
)

func TimeOuts(){
	c1:=make(chan string,1)
	c2:=make(chan string,1)

	go func(){
		time.Sleep(2*time.Second)
		c1 <- "result one"
	}()

	select{
	case res:= <-c1:
		fmt.Println(res)
	case <- time.After(1*time.Second):
		fmt.Println("time out 1")
	}

	go func(){
		time.Sleep(2*time.Second)
		c2 <- "result two"
	}()

	select{
	case res:= <-c2:
		fmt.Println(res)
	case <- time.After(3*time.Second):
		fmt.Println("time out 2")
	}


}