package main
import "fmt"

func main() {
	var sum = 0
	for i:=1;i<11;i++ { // Normal For loop
         sum+=i
	}
	fmt.Println(sum)
    
	for { // Infinite for loop
		break
	}
    i:=0
	for i < 5 {
		// Conditional for loop
		i++ 
	}

	// for loop with goto 
	for {
         if i < 6 {
			goto end
		 }
		 fmt.Println("Normal")
	}
	end:
	fmt.Println("End")


}