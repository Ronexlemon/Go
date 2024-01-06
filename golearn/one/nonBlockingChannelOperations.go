package main

import (
	"fmt"
)

func NonBlockingChannelOperation(){
	messages := make(chan string)
	signal:=make(chan string)

	select{
	case msg:= <-messages:
		fmt.Println("received message",msg)
	default:
		fmt.Println("no message Received")
	}

	msg:="hi"

	select{
	case messages <-msg:
		fmt.Println("sent message",msg)
	default:
		fmt.Println("no message sent")
	}

	select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signal:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }


}