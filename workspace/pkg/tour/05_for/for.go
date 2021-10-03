// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

package main

import "fmt"

func optionalStatements(){
	// The init and post statements are optional.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}


func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	optionalStatements()
}
