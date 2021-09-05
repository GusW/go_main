package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(sli []int, i int) {
	tmp := sli[i+1]
	sli[i+1] = sli[i]
	sli[i] = tmp
}

func BubbleSort(sli []int) {
	for {
		sortedFlag := true
		for i := 0; i+1 < len(sli); i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				sortedFlag = false
			}
		}
		if sortedFlag == true {
			return
		}
	}
}

func getUserInput() string {
	var userInput string
	/* We're using bufio Scan because fmt.Scan will fail
	if user will try to enter more than one string
	*/
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}
	return userInput
}

func convertStringToIntSlice(s string) []int {
	strings := strings.Split(s, " ")

	intSlice := make([]int, len(strings))
	var err error
	for i, s := range strings {
		intSlice[i], err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Fatal error!", err)
			os.Exit(1)
		}
	}
	return intSlice
}

func main() {
	fmt.Println("Welcome to Bubble Sort!")
	fmt.Println("You will be asked to input the array of integers (use space as divider)")
	fmt.Println("Example format: -1 1 4 0 5 9")
	fmt.Print("Please enter: ")

	userInput := getUserInput()
	intSlice := convertStringToIntSlice(userInput)

	BubbleSort(intSlice)
	for _, v := range intSlice {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
