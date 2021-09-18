package main

import "fmt"

func main() {

	/*
		Race condition is when multiple threads(go routines) are trying to execute pieces of code without
		any kind of synchronism and there are some kind of dependency between or expected result order.

		Another way is when they access and manipulate the same variable.

		In the code below, due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.

		There is absolute no garantee in the order this strings will be displayed because the go scheduler can mess it up

	*/
	go fmt.Printf(" ||| In your marks... |||")
	go fmt.Printf(" ||| Go |||")
	for i := 0; i < 200; i++ {
		fmt.Printf(" %d ", i)
	}

}
