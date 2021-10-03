package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1<<6)
	fmt.Println("My favorite number is", rand.Intn(9))
}
