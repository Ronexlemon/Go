package main

import (
	"fmt"
)

func ping(pings chan<- string,msg string){
	pings <- msg;
}
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func ChannelDirection(){
	pings:=make(chan string, 1 )
	pongs :=make(chan string,1)

	ping(pings,"from this channel")
	pong(pings,pongs)

	fmt.Println(<-pongs)
}