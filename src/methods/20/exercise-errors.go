package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("No square root for negative numbers: %v", float64(*err))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)

		return 0, &err
	}

	z := x / 2.0

	for last := 0.0; math.Abs(last-z) >= 0.00000000001; {
		last = z

		z -= (z*z - x) / (2 * z)

		fmt.Printf("z -> %v\n", z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
