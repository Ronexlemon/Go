package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounter(){
	// Create an atomic counter ops
	var ops atomic.Uint64
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
// Launch 50 goroutines
	for i:=0;i< 50;i++{
		// Increment the WaitGroup counter for each goroutine launched
		wg.Add(1)
		// Launch a goroutine for each iteration
		go func(){
			for c:=0;c<1000;c++{
				// Perform 1000 increments on the ops counter
				ops.Add(1)
			}
			// Decrement the WaitGroup counter when the goroutine finishes
			wg.Done()

		}()
	}
	// Wait for all goroutines to finish before exiting the main function
	wg.Wait()
	fmt.Println("ops:",ops.Load())
}