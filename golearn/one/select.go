package main

import (
	"fmt"
	"time"
)
/*select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go. */
func channel1(ch chan<- string,msg string){
	time.Sleep(time.Second)
	ch <- msg
}
func channel2(ch2 chan<-string,msg string){
	time.Sleep(time.Second)
	ch2<- msg
}
func Select(){
	c1:=make(chan string)
	c2:=make(chan string)

	go channel1(c1,"one")
	go channel2(c2,"two")

	for i:=0;i<2;i++{
		select{
		case msg1:= <-c1:
			fmt.Println("received",msg1)
		case msg2:=<-c2:
			fmt.Println("received",msg2)
		}
	}
}