package main

import (
	"fmt"
	mathematic "golang-book-tasks/chapter-11/math"
)

func main() {
	slice := []float64{231.213, 123.213, 4522.2, 2.44}
	fmt.Println("Our slice: ", slice)
	fmt.Println("Average value: ", mathematic.Average([]float64{}))
	fmt.Println("Max value: ", mathematic.Max(slice))
	fmt.Println("Min value: ", mathematic.Min(slice))
}
