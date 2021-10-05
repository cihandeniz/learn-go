package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, &MyError{time.Now(), "divide by zero"}
	}

	return x / y, nil
}

func main() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
