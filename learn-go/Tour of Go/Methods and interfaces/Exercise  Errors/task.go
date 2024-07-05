package main

import (
	"fmt"
	"math"
)

/* TODO */

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", e)
}

//func Sqrt(f float64) (float64, error) {
//	if f >= 0 {
//		return math.Sqrt(f), nil
//	} else {
//		return 0, ErrNegativeSqrt(f)
//	}
//}

func Sqrt(f float64) (float64, error) {
	if f >= 0 {
		const delta = 1e-10
		z := f
		for {
			n := z - (z*z-f)/(2*z)
			if math.Abs(n-z) < delta {
				break
			}
			z = n
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(f)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
