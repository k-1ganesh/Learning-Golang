package main 
import "fmt" 

func main() {
	var arr = [5]int{1,2,3,3,4}
	fmt.Println(arr)

	ar := [3]string{"a" , "b" , "c"}
	fmt.Println(ar)
	// Lenght of array 
	var lenght = len(arr)
	fmt.Println(lenght)

	// Iterating through array.
	for i:=0;i<lenght;i++ {
		fmt.Print(arr[i])
	}

	// Iterating using range.
	for index , value := range arr {
		fmt.Println(index , value)
	}

}