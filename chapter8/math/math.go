package math

// calcula a média de uma série de números
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	min := 0.0
	for i, v := range xs {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

func Max(xs []float64) float64 {
	max := 0.0
	for i, v := range xs {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}
