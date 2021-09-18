package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("")
	fmt.Println("What are race conditions?")
	fmt.Println("")
	fmt.Println("Race conditions occur when the program outcome depends on a non-deterministic interleaving between processes/threads exchanging state.")
    fmt.Println("Once concurrenct execution depends primarily in how the base OS handles process scheduling -")
    fmt.Println("giving its own scheduling algorithm - processes'context switch is not deterministic and mutliple possible")
    fmt.Println("outcomes can be derived from the race conditions.")
	fmt.Println("")
	fmt.Println("A working Goroutine example with race conditions can be triggered below!")
	fmt.Println("Let the program running for a minute or so to notice the non-deterministic outputs caused by race conditions:")
	fmt.Println("")

	for {
		var targetNumber float64
		targetNumber = 3.45
		random := rand.Intn(3)
		go func() {
			targetNumber ++
			time.Sleep(time.Duration(random) * time.Second)
			targetNumber =+ 9.87
			targetNumber = math.Pow(targetNumber, 1.23)
			targetNumber = math.Pow(targetNumber, 2)
			targetNumber -= 14
			targetNumber = math.Pow(targetNumber, 9)
		}()

		go func() {
			targetNumber *= 6.78
			time.Sleep(time.Duration(random) * time.Second)
			targetNumber = math.Sqrt(targetNumber)
		}()

		time.Sleep(1 * time.Second)
		fmt.Println("MAIN =>", targetNumber)
		fmt.Println("")
	}
}
