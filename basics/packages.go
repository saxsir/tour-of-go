package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Added
	fmt.Println("My favorite number is", rand.Intn(10))
}