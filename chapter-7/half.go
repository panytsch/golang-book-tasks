package main

import (
	"fmt"
	"os"
)

func main() {
	var input int
	fmt.Print("Write a integer number: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		os.Exit(0)
	}
	i, b := half(input)
	fmt.Println("result: ", i, b)
}

func half(number int) (result1 int, result2 bool) {
	result1 = number / 2
	result2 = result1%2 == 0
	return
}
