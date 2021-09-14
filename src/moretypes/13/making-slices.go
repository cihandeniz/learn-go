package main

import "fmt"

func main() {
	len := 5
	slice := 2

	a := make([]int, len)
	printSlice("a", a)

	b := make([]int, 0, len)
	printSlice("b", b)

	c := b[:slice]
	printSlice("c", c)

	d := c[slice:len]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
