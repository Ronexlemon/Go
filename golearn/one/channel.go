package main

import (
	"fmt"
)

//channels are pipes thatconnect concurrent goroutines.

func makeChannel(){
	message:= make(chan string)
	//sending a value to a channel using syntax chanel <- 
	go func(){message <- "sending to another goroutine"}()

	//receive the message
	msg:= <- message

	fmt.Println(msg)

}

func Channel(){
	makeChannel()
}