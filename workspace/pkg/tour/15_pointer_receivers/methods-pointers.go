// Pointer receivers
// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
// For example, the Scale method here is defined on *Vertex.
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// With a value receiver, the Scale method operates on a copy of the original Vertex value.
// (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
// There are two reasons to use a pointer receiver.
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// Here we see the Abs and Scale methods rewritten as functions.
// Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer
// while methods with pointer receivers take either a value or a pointer as the receiver when they are called

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Scale(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 10)
// 	fmt.Println(Abs(v))
// }
