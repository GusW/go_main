package main

import "fmt"

func main(){
	var value float64
	fmt.Println("Please write me a float number")
	fmt.Scanf("%f", &value)
	fmt.Printf("You have entered the value: %.0f", value)
}

