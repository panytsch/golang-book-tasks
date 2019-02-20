package main

import (
	"fmt"
	"os"
)

// cached values
var fibSlice = map[uint8]uint64{0: 0, 1: 1}

func main() {
	var input uint8
	fmt.Print("Write a number of fibonacci numbers: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		os.Exit(0)
	}
	fmt.Println("Your fibonacci number is: ", fibonacci(input))
}

func fibonacci(n uint8) uint64 {
	if n == 0 {
		return 0
	} else if val, exist := fibSlice[n]; exist {
		return val
	}
	fibSlice[n] = fibonacci(n-1) + fibonacci(n-2)
	return fibSlice[n]
}
