package main

import (
	"fmt"
	"time"
)

/*
tickers are for when you want to do something repeatedly at regular intervals
*/


func Tickers(){
	ticker:= time.NewTicker(500 * time.Millisecond)
	done :=make(chan bool)

	go func(){
		for{
			select{
			case <-done:
				return
			case t:=<-ticker.C:
				fmt.Println("Tick At",t)
			}
		}
		
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <-true
	fmt.Println("Ticker stopped")
}