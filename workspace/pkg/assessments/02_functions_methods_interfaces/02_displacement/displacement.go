package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64{
		return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
}

func _hanleInputOf(label string) float64{
	var variable float64
	prefixErr := "Could not capture"
	fmt.Println("Enter a value for", label,"and press Enter:")
	_, err := fmt.Scan(&variable)
	if err != nil {
		fmt.Println(prefixErr, label, ": ", err)
		panic("Please restart application and try again")
	}
	return variable
}

func main(){
	accelerationLabel := "acceleration"
	initialVelocityLabel := "initial velocity"
	initialDisplacementLabel := "initial displacement"
	timeLabel := "time"

	acceleration := _hanleInputOf(accelerationLabel)
	initialVelocity := _hanleInputOf(initialVelocityLabel)
	initialDisplacement := _hanleInputOf(initialDisplacementLabel)
	partialDisplacementFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	time := _hanleInputOf(timeLabel)
	totalDisplacement:= partialDisplacementFn(time)
	fmt.Println("Total displacement is: ", totalDisplacement)
}
