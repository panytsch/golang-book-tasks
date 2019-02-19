package main

import "fmt"

//C = (F - 32) * 5/9
func main() {
	fmt.Print("Write temperature in Fahrenheit: ")
	var temperature float64
	_, _ = fmt.Scanf("%f", &temperature)
	fmt.Printf("Temperature in Celsius is %.2f\n", convertFtoC(temperature))
}

func convertFtoC(t float64) float64 {
	return (t - 32) * 5 / 9
}
