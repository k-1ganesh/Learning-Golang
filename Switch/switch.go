package main
import "fmt"

func main() {
	var a = 1
	switch a {  // In GO only the matched case is executed so no need to break
	case 1: 
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Wrong")
	}


	// Multicase Switch
	switch a {
	case 1,2:
		fmt.Println("One or Two")
	case 3,4:
		fmt.Println("Three or Four")
	}
}