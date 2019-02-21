package math

//find average value if float64 slice
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//find max value of float64 slice
func Max(slice []float64) (max float64) {
	max = slice[0]
	for _, i := range slice[1:] {
		if i > max {
			max = i
		}
	}
	return
}

//find min value of float64 slice
func Min(slice []float64) (min float64) {
	min = slice[0]
	for _, i := range slice[1:] {
		if i < min {
			min = i
		}
	}
	return
}
