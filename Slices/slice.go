// Slice is a sliced part of array. 
// Slice can grow as needed.
// Slice created from array arn't actual instead they are just reference to that array.
package main 
import "fmt" 

func main() {
	slice1 := []int{1,2,3,4} // This is one way of making slice.
	fmt.Println(slice1)
    array := [5]int{1,2,3,4,5} // This is array of fixed size.
	slice2 := array[1:4] // This is another way of defining slice.
	fmt.Println(slice2)
    slice3 := make([]int , 3, 5) // 3 -> Length , 5-> Capacity.
	fmt.Println(slice3) 



	// Appending element to the slice.
    slice1 = append(slice1 , 4 , 5 , 6) 
	fmt.Println(slice1) 

	slice1 = append(slice1 , slice2...)  
	fmt.Println(slice1) 
    
}