package main

import (
	"fmt"
)

type DivisionError struct {
	Dividend float64
	Divisor  float64
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %f by %f", e.Dividend, e.Divisor)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
