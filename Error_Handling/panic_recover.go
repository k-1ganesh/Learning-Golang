// Panic directly crash the program.
// What happens when panic occurs ?
// 1) Program execution stops.
// 2) Defer statements gets executed.
// A -> B -> C Suppose this is function call and panic occurs at C then 1 & 2 statement will happen and same thing will happen to B and A aswell untill control reaches main and program terminates.

package main

import "fmt"

func fun2() {
	fmt.Println("Start fun2")
	defer fmt.Println("End fun2")
	panic("Panic Occured")
	fmt.Println("Bye")
}
func fun1() {
	fmt.Println("Start fun1")
	defer fmt.Println("End fun1")
	fun2()
	fmt.Println("HI")
}
func main() {
	fmt.Println("Hello World")
	fun1()
	fmt.Println("Bye World")
}
