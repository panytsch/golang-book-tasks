package main

import (
	"fmt"
	"os"
)

func main () {
	var input1, input2 float64
	fmt.Print("Write first and second number separated by space: ")
	_, err := fmt.Scanf("%f %f", &input1, &input2)
	if err != nil {
		fmt.Println("Wrong!")
		os.Exit(0)
	}
	fmt.Println("first number = ", input1, "\nSecond number = ", input2)
	swap(&input1, &input2)
	fmt.Println("Numbers after swap:\nfirst = ", input1, "\nsecond = ", input2)
}

func swap (a *float64, b *float64) {
	*a, *b = *b, *a
}