package main

import "fmt"

func add(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(add(42, 13, 800))
}
