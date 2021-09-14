package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2.0

	for last := 0.0; math.Abs(last-z) >= 0.00000000001; {
		last = z

		z -= (z*z - x) / (2 * z)

		fmt.Printf("z -> %v\n", z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(5))
	fmt.Println(Sqrt(6))
	fmt.Println(Sqrt(7))
	fmt.Println(Sqrt(8))
	fmt.Println(Sqrt(9))
}
