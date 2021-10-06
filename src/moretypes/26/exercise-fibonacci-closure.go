package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cur, last := -1, -1

	return func() int {
		if last < 0 {
			last = 0
			return last
		} else if cur < 0 {
			cur = 1
			return cur
		}

		last, cur = cur, cur+last

		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 25; i++ {
		fmt.Println(f())
	}
}
