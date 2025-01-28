// When you have multiple goroutines, you might want to wait for all of them to finish before proceeding. This is where sync.WaitGroup comes in.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started goroutine id : %v\n", id)
	time.Sleep(time.Second * 1)
	fmt.Printf("Ending goroutine id : %v\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Closing Main....")
}
