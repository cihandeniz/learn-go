package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Printf("i: %v, sum: %v\n", i, sum)
	}

	fmt.Printf("\nfinal sum: %v\n", sum)
}
