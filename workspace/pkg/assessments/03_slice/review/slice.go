package main

import (
	"fmt"
	"sort"
)

func main() {
	sortme := make([]int, 4)

	for i := 0; i < len(sortme); i++ {
		fmt.Println("Enter a integer or 'X' to exit: ")
		_, err := fmt.Scanf("%d", &sortme[i])

		if err != nil {
			fmt.Println("Int could not be read. Exiting.", err)
			return
		}

		tmp_slice := sortme[:i+1]
		sort.Ints(tmp_slice)
		fmt.Println(tmp_slice)
	}
}
