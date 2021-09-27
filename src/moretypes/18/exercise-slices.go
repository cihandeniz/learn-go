package main

import (
	"math"
	// "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = make([]uint8, dx)

		for x := range result[y] {
			result[y][x] = uint8(math.Abs(float64(dx/2-x)) + math.Abs(float64(dy/2-y)))
		}
	}

	return result
}

func main() {
	// pic.Show(Pic)
}
