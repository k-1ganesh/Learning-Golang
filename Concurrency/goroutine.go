package main
// Go routins are the threads which runs councurrently. 
// Use go keyword to make go routine
import (
	"fmt"
	"time"
)

func numPrint(num int) {
    fmt.Println(num)
}

func main() {
     go numPrint(2)
     go numPrint(3)
     go numPrint(1)
     
     time.Sleep(time.Second*1)
     fmt.Println("Bye")
}