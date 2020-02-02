package main

import (
	"fmt"
)

func main() {
	test(split(20))
	test(swap("cihan", "deniz"))
	test(add(2, 3))
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
