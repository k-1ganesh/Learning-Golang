package main 
import "fmt" 

// Syntax Of Function 
// func Name(Parameter1,...) Return type { }
// func add(a int , b int) int { return a + b}

func add(a int , b int) int {
	return a + b
}
func concat(s1 string , s2 string) string {
	return s1 + s2
}



// Function Returning Multiple value.
func AddSub(a int , b int) (int , int) {
	return a+b , a-b
}



// Variadic Function 
// A function which can have variable no. of parameters.
func Addition(numbers ...int) int {
	sum := 0
	for _,num := range numbers {
		sum += num 
	}
	return sum
}

// Recursive Function 
func Sum(num int) int {
	if num == 0 {
		return 0
	} else {
		return num + Sum (num - 1)
	}
}

// Assigning function to Variable.
func performOperation(a int , b int , operation func(int , int) int) int {
	return operation(a , b)
}

func main() {
     fmt.Println(add(10 , 15))
	 fmt.Println(concat("Ganesh" , "Kachare"))
	 add,sub := AddSub(10,15)
	 fmt.Println(add,sub)
	 // _ is used to ignore the value. 
	 // add , _ := AddSub(10,15)
	 fmt.Println(Addition(2,3,4,5,6,7))
	 fmt.Println(Sum(10))

	 funVariable := func(a int , b int ) int {return a * b} // Assigned function to variable.
	 fmt.Println(performOperation(10 , 20 , funVariable))
}