package main

import "fmt"

func main() {
	fmt.Print("Write feet number: ")
	var feet float64
	_, _ = fmt.Scanf("%f", &feet)
	fmt.Printf("Meter value is %.2f\n", convertFeetToMeter(feet))
}

func convertFeetToMeter(feet float64) float64 {
	return feet * 0.3048
}
