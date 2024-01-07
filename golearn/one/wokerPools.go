package main

import (
	"fmt"
	"time"
)

func workerspool(id int, jobss <-chan int, results chan<- int) {
	for j := range jobss {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPool() {
	const numJobs = 5
	jobss := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workerspool(w, jobss, results)
	}
	for j:=1;j <=numJobs;j++{
		jobss <-j
	}
	close(jobss)

	for a:=1;a<= numJobs;a++{
		<-results
	}
}
