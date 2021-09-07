package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	val := 100
	rand.Seed(time.Now().Local().Unix())
	for val > 9 {
		val = rand.Intn(100)
		fmt.Println("rand num", val)
	}
}
