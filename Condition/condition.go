package main
import "fmt"

func main() {
	var x = 10
	var y = 15

	// If else Statement.
	if (x > y) {
		fmt.Println("X is greater")
	} else {   // else should be like this only } else {  Else error will be generated.
		fmt.Println("X is smaller")
	}

	// Else if statement.
	var a = 3
	var b = 4
	var c = 5
	if (a > b && a>c) {
        fmt.Printf("%v is max",a)
	} else if (b > c) {
		fmt.Printf("%v is max",b)
	} else{
		fmt.Printf("%v is max" , c)
	}
}	