package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0 // all goroutines will try to modify this counter simultaneously

	const num = 2 // 2 self-executing goroutine functions will be launched
	var wg sync.WaitGroup  // To wait for multiple goroutines to finish, we can use a wait group.
	wg.Add(num)    // The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished.


	// Mutex can be used to resolve the race condition. Lock() the variable just before reading it and Unlock() is just after writing to the variable.
	//var mu sync.Mutex

	for i := 0; i < num; i++ {
		go func() {     // each func reads the value of counter, increments it and writes back
	
	//		mu.Lock()
	
			temp := counter
			runtime.Gosched()  // Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			temp++
			counter = temp
	
	//		mu.Unlock()
	
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
