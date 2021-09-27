package main

import (
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, word := range strings.Fields(s) {
		result[word] = result[word] + 1
	}

	return result
}

func main() {
	// wc.Test(WordCount)
}
