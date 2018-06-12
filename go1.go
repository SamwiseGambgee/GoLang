package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("addMe function call returns an Integer: ", addMe(2, 2))
	fmt.Println("The value of Pi is: ", math.pi)
}

func addMe(a, b int) int {
	return a + b
}
