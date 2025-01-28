//When multiple goroutines access shared data, you need to prevent race conditions. A sync.Mutex ensures that only one goroutine can access the data at a time.

package main 

import (
	"fmt" 
	"sync"
)
 
var counter int = 0  // Global variable.
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done() 
    mutex.Lock()
	counter++  // Critical Section.
	mutex.Unlock()
}

func main() {
    
	var wg sync.WaitGroup 
	for i := 1;i<=10;i++ {
		wg.Add(1)
		increment(&wg)
	}
	wg.Wait()
   fmt.Println("Ending Main...")
   fmt.Println(counter)
}