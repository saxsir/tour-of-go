package main

import (
	// "fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)

	for _, w := range f {
		if c, e := m[w]; e == true {
			m[w] = c + 1
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
