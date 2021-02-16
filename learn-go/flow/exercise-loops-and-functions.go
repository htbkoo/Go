package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z = float64(1)

	const MAX_ITERATIONS = 10
	for i := 1; i <= MAX_ITERATIONS; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
