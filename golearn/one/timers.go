package main

import (
	"fmt"
	"time"
)

/*
We often want to execute Go code at some point in the future, or repeatedly at some interval. Go’s built-in timer and ticker features make both of these tasks easy
*/
func Timers(){
	timer1:= time.NewTimer(2* time.Second)

	<-timer1.C
	fmt.Println("Timer 1 is fired")
	timer2:= time.NewTimer( time.Second)

	go func(){
		<-  timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2:=timer2.Stop()
	if stop2{
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2*time.Second)
}

