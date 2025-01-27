package main

import "fmt"

// If panic is resolved then parent function will not halt.
func fun1() {
	fmt.Println("Start of fun1")
	defer fmt.Println("End of fun1")
	defer func() { // WIth the help of ananomous function and defer panic is recovered.
		if r := recover(); r != nil {
			fmt.Println(r, "Panic Resolved")
		}
	}()
	panic("Panic Occured")
}

func main() {
	fmt.Println("Hello !!")
	fun1()
	fmt.Println("Bye !!")
}
