package main

import (
	"fmt"
	"math"
)

var totalLoopCount int

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		totalLoopCount++

		oldZ := z
		z = z - (math.Pow(z, 2)-x)/(2*z)

		if (z - oldZ) < 0.0000000001 {
			break
		}
	}

	return z
}

func main() {
	totalLoopCount = 0

	for i := 1; i <= 10; i++ {
		a := Sqrt(float64(i))
		b := math.Sqrt(float64(i))
		if a == b {
			fmt.Println("ok")
		} else {
			fmt.Printf("sqrt(%v): %v\n", i, a)
			fmt.Printf("math.sqrt(%v): %v\n", i, b)
		}
	}

	fmt.Printf("Total loop count is %v", totalLoopCount)
}
