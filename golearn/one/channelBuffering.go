package main

import (
	"fmt"
)

/*
By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.*/

func channelBuffering(){
	message:=make(chan string,4)
	message <- "buffered"
	message <- "channel"
	

	fmt.Println(<-message)
	fmt.Println(<-message)
	
	
}

func ChannelBuffered(){
	channelBuffering()
}