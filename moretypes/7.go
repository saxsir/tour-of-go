package main

import "fmt"

func main() {
	// slices
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	fmt.Println("#####")

	// array
	var a [2]int
	a[0] = 0
	a[1] = 1
	fmt.Println("a ==", a)

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] == %d\n", i, a[i])
	}
}
