package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("t is <nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	if i == nil {
		fmt.Println("i is <nil>")
	}

	var t *T
	i = t

	if i != nil {
		fmt.Println("i is not <nil>")
	}

	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
