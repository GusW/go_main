// Pointers to structs
// Struct fields can be accessed through a struct pointer.
// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)
}
