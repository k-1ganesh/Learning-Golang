// Pointers are used for the reference purpose.
// Pointer arithmetic is not supported in Golang. 
package main 
import "fmt" 


type reactangle struct {
	lenght float64
	width float64
}

func area(rect *reactangle) float64 { // Using pointer in function. 
	return *&rect.lenght *  *&rect.width
}

func swap(a *int , b *int) {
	temp := *a
	*a = *b
	*b = temp 
}

func main() {
	num := 100
	ptr1 := &num  // First way of defining and assining value.
    var ptr2 *int = &num  // Second way of defining and assigning value.
	fmt.Println(ptr1) 
	fmt.Println(*ptr1)
	fmt.Println(ptr2) 
	fmt.Println(*ptr2)
	rect1 := reactangle{10,20}
	fmt.Println(area(&rect1))
	a := 10 
	b := 15 

	swap(&a , &b) 
	fmt.Printf("a : %v , b: %v",a , b)
}