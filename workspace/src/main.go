package main

import "fmt"

type P struct {
	x string
	y int
}

func main() {
	b := P{"x", -1}
	a := [...]P{{"a", 10}, {"b", 2}, {"c", 3}}
	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)
}
