package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cur, last := 0, 0

	return func() int {
		if cur == 0 && last == 0 {
			cur = 1
		} else {
			cur = cur + last
			last = cur - last
		}

		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 25; i++ {
		fmt.Println(f())
	}
}
