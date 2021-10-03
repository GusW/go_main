// Array
// An array's length is part of its type, so arrays cannot be resized.
// Slice
// A slice does not store any data, it just describes a section of an underlying array.
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// Other slices that share the same underlying array will see those changes.
// A slice literal is like an array literal without the length.
// A slice has both a length and a capacity.
// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s)
// The zero value of a slice is nil.
// A nil slice has a length and capacity of 0 and has no underlying array.
// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
// Slices can contain any type, including other slices.
package main

import "fmt"

func slicing(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiteral(){
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func nilSlice(){
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func dynamicSlice(){
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func appending(){
	var s []int
	printSlice("s0", s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice("s1", s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice("s2", s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice("s3", s)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	slicing()

	sliceLiteral()

	nilSlice()

	dynamicSlice()

	appending()
}
