package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := -1
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", f())
	}
}
