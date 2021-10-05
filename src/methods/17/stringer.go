package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Stringer interface {
	String() string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func print(values ...interface{}) {
	for i, value := range values {
		if i > 0 {
			if i == len(values)-1 {
				fmt.Print(" and ")
			} else {
				fmt.Print(", ")
			}
		}

		switch stringer := value.(type) {
		case Stringer:
			fmt.Print(stringer.String())
		default:
			fmt.Printf("%v (with default formatter)", stringer)
		}
	}

	fmt.Println()
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	print(a, z, "me")
}
