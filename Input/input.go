package main

import (
	"bufio" // Used to read entire line
	"fmt"
	"os"
	"reflect" // Used to check the type  of Variable. reflect.TypeOf
	"strconv" // Used to convert string into something.
	"strings" // Used to handle string operations.
)



func main() {

	var age int
	fmt.Println("Enter the Age:")
	fmt.Scanln(&age) // Single Input.
	if age > 18 {
		fmt.Println("Eligible for voting !!")
	} else {
		fmt.Println("Not Eligible for voting !!")
	}

	// Multiple Input. 
	var a int
	var b int
	fmt.Println("Enter the value of a & b") 
	fmt.Scan(&a , &b) 
	if a>b{
		fmt.Println("A is greater")
	} else {
		fmt.Println("B is greater")
	}

	// Multiple input Comma seperated.
	// To do this 2 new imports are required bufio to read line and strings to handle string operations.
    
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Principle Rate And Time")
	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	values := strings.Split(input , ",")
	fmt.Println(reflect.TypeOf(values))

	principle , _ := strconv.ParseFloat(values[0] , 64)
	rate , _ := strconv.ParseFloat(values[1],64)
	time , _ := strconv.ParseFloat(values[2],64)
	
	Compound_Interest := principle * rate * time / 100
    
	fmt.Printf("Compound Interest is %v",Compound_Interest)

}