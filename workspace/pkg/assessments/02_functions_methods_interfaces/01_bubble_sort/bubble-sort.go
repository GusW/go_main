package main

import "fmt"


func Swap(sequence []int, idx int){
	// should return nothing
	// should swap the contents of the slice in position i with the contents in position i+1.
	currentValue := sequence[idx]
	sequence[idx] = sequence[idx+1]
	sequence[idx+1] = currentValue
}

func BubbleSort(sequence []int) {
	// returns nothing
	// should modify the slice so that the elements are in sorted order
	for idx, value := range(sequence){
		nextTargetIdx := idx + 1
		if nextTargetIdx == len(sequence){
			break
		}
		if value > sequence[nextTargetIdx]{
			Swap(sequence, idx)
			BubbleSort(sequence)
		}
	}
}

func main(){
	userSequence := make([]int, 0, 10)
	for len(userSequence) < cap(userSequence) {
		var currentValue int
		fmt.Print("\nEnter the integer #", len(userSequence) + 1, " of the sequence and press Enter:\n")
		_, err := fmt.Scan(&currentValue)
		if err != nil {
			fmt.Print("\nCould not capture integer: ", err, "\nPlease try again or 'Ctrl + C' to abort.\n")
		} else {
			userSequence = append(userSequence, currentValue)
		}
	}
	BubbleSort(userSequence)
	fmt.Print("Bubble sorted sequence is: ", userSequence, "\n")
}
