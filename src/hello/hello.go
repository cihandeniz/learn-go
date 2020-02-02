package main

import (
	"fmt"
	"math"
)

func main() {
	// i'm here -> https://tour.golang.org/flowcontrol/1

	test(numericConstants())
	test(constants("world"))
	test(typeInference())
	test(conversions(3, 4))
	test(zeroValues())
	test(basicTypes())
	test(vars())
	test(split(20))
	test(swap("right", "left"))
	test(add(2, 3))
}

const (
	// NumericConstantsBig is 1 with 100 following 0s
	NumericConstantsBig = 1 << 100

	// NumericConstantsSmall is just 10
	NumericConstantsSmall = NumericConstantsBig >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func numericConstants() (int, float64, float64) {
	return needInt(NumericConstantsSmall),
		needFloat(NumericConstantsSmall),
		// needInt(NumericConstantsBig), // doesn't compile
		needFloat(NumericConstantsBig)
}

func constants(world string) (string, string) {
	const Hello = "hello"

	return Hello, world
}

func typeInference() complex128 {
	x := 1.2 + 2i

	return x
}

func conversions(x, y int) (int, int, uint) {
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	return x, y, z
}

func zeroValues() (i int, f float64, b bool, s string) {
	return
}

var (
	basicTypesToBe   bool   = false
	basicTypesMaxInt uint64 = 1<<64 - 1
)

func basicTypes() (bool, uint64) {
	return basicTypesToBe, basicTypesMaxInt
}

var varsX, varsY, varZ = 1, "two", true

func vars() (int, string, bool, int) {
	i := 4

	return varsX, varsY, varZ, i
}

func split(sum int) (x, y int) {
	x = sum * 2 / 9
	y = sum - x

	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

func test(uut ...interface{}) {
	fmt.Println("result is:", uut)
}
