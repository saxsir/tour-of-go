package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	var f = math.Sqrt(float64(x*x + y*y))
	var z = f
	fmt.Println(x, y, z)
}
