package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) { // error is a interface
	if b == 0 {
		return 0, errors.New("Divide by zero Error") // Used to create a simple error msg.
	}
	return a / b, nil
}

// There is another way to format and return error.
// fmt.Errorf("formated error...") Can be used instead of errors.New()
func main() {
	ans, err := division(10, 0)
	if err != nil {
		fmt.Println("Error Occured:", err)
	}
	fmt.Println(ans)
}
