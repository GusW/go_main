// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	res := make(map[string]int)
	for _, str := range strs {
		switch _, ok := res[str]; ok {
		case true:
			res[str] += 1
		case false:
			res[str] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
