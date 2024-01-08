package main

import (
	"fmt"
	"sync"
	"time"
)

//To wait for multiple goroutines to finish, we can use a wait group

//function to run in every goroutine

func Work(id int){
	fmt.Printf("Worker %d starting\n",id)
	time.Sleep(time.Second)//sleep to simulate an expensive task
	fmt.Printf("worker %d done\n",id)

}

func WaitGroup(){
	var wg sync.WaitGroup
	for i:=1;i <=5; i++{
		wg.Add(1)
		i:=i

		go func(){
			defer  wg.Done()
			Work(i)
		}()

	}
	wg.Wait()
}