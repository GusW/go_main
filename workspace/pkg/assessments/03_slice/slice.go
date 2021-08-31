package main

import (
	"fmt"
	"strconv"
)

func _handleUserInput() (string, bool) {
	var currentInput string
	fmt.Print("Enter any Integer or 'X' to exit app:\n")
	_, err := fmt.Scan(&currentInput)
	if err != nil {
		fmt.Printf("There was an error while capturing the number: %s\n", err)
		return "continue", false
	}
	if currentInput == "X" {
		return "break", false
	}

	return currentInput, true
}

func _handleSortedSlice(currentSlice []int64, newElement string) []int64 {
	newInt, err := strconv.ParseInt(newElement, 10, 64)
	if err != nil {
		fmt.Printf("There was an error while parsing the number: %s\n", err)
		return nil
	} else {
		for idx, val := range currentSlice {
			if newInt <= val {
				newSlice := []int64{}
				newSlice = append(newSlice, currentSlice[0:idx]...)
				newSlice = append(newSlice, newInt)
				newSlice = append(newSlice, currentSlice[idx:]...)
				return newSlice
			}
		}
		currentSlice = append(currentSlice, newInt)
		return currentSlice
	}
}

func main() {
	userNumbers := make([]int64, 3)
	userNumbers = userNumbers[:0]
	for {
		userInput, isSucceeded := _handleUserInput()
		if !isSucceeded {
			if userInput == "continue" {
				continue
			}
			break
		}
		userNumbers = _handleSortedSlice(userNumbers, userInput)
		if userNumbers != nil {
			fmt.Printf("%v\n", userNumbers)
		} else {
			fmt.Print("Not an int, punk...\n")
		}
	}
}
