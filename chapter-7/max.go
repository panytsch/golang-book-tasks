package main

import "fmt"

func main() {
	a := 1.0
	b := 45.3
	c := -5.1
	d := 3.0
	f := 12.00005
	fmt.Println("Start values is: ", a, b, c, d, f)
	fmt.Println("Maximum number is: ", maximum(a, b, c, d, f))
}

func maximum(params ...float64) float64 {
	max := params[0]
	for _, i := range params {
		if i > max {
			max = i
		}
	}
	return max
}
