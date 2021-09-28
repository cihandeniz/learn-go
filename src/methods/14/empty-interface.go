package main

import "fmt"

type any interface{}

func main() {
	var i any
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
