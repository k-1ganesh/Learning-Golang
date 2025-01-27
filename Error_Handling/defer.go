// Defer keyword is used to delay the execution of statment till the end of function.
// Defer statments are executed at the end of function no matter even error occurs.
// defer statments uses defer stack for this which follows lifo order.
// defer can be used for closing resources.
package main  
import (
	"fmt"
)

func main() {
    fmt.Println("File Opening") 
	defer fmt.Println("File Closing")
	fmt.Println("1")
	fmt.Println("2")
	panic("Panic Occured.")
	fmt.Println("3") // This 2 lines will not execute after panic. But defer statements will be executed.
	defer fmt.Println("Last End")
}