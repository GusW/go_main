package main

import "fmt"

func main() {
	var acceleration float64
	fmt.Print("Please enter the acceleration:   ")
	fmt.Scanln(&acceleration)

	var initial_velocity float64
	fmt.Print("Please enter the initial velocity:  ")
	fmt.Scanln(&initial_velocity)

	var initial_displacement float64
	fmt.Print("Please enter the initial displacement:  ")
	fmt.Scanln(&initial_displacement)

	var time float64
	fmt.Print("Please enter time:  ")
	fmt.Scanln(&time)

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Println(fn(3))
	fmt.Println(fn(5))


}


func GenDisplaceFn(acceleration, initial_velocity, initial_displacement float64) func(float64) float64{
	return func(time float64) float64 {
		return 1.0/2.0 * acceleration * time * time + initial_velocity * time + initial_displacement
	}
}
