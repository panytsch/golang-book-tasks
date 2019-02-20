package main

import "fmt"

func main() {
	var sliceFloat []float64
	for {
		var input float64
		fmt.Print("Write a number, or something wrong to exit: ")
		_, err := fmt.Scanf("%f", &input)
		if err != nil {
			break
		}
		sliceFloat = append(sliceFloat, input)
	}
	fmt.Println(sum(sliceFloat))
}

func sum(a []float64) float64 {
	result := 0.0
	for _, i := range a {
		result += i
	}
	return result
}
