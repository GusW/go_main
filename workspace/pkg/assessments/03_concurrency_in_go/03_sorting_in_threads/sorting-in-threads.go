package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)


func sortNumberSlice(userNumbersSlice []int) []int {
	sort.Ints(userNumbersSlice)
	return userNumbersSlice
}

func sortRoutine(sortChannel chan []int, userNumbersSlice []int){
	fmt.Println("Sorting slice", userNumbersSlice, "...")
	sortChannel <- sortNumberSlice(userNumbersSlice)
}

func handleSort(userNumbers []int, subProcessesAmount int){
	sortChannel := make(chan []int, subProcessesAmount)
	userNumbersLen := len(userNumbers)
	subSliceRate := float64(len(userNumbers))/float64(subProcessesAmount)
	intSubSlice, floatSubSlice := math.Modf(subSliceRate)
	startIdx := 0
	var endIdx int
	for i := 0; i < subProcessesAmount; i ++{
		nextIterIdx := i + 1
		if nextIterIdx == subProcessesAmount {
			endIdx = userNumbersLen
		} else {
			endIdx = startIdx + int(intSubSlice)
			if floatSubSlice == 0.5 && nextIterIdx > int(intSubSlice) && nextIterIdx <= userNumbersLen{
				endIdx += 1
			} else if floatSubSlice != 0{
				if floatSubSlice < 0.5 {
					if intSubSlice > 1{
						if nextIterIdx == 1 {
							endIdx += 1
						}
					} else if nextIterIdx < int(intSubSlice) && endIdx > 0{
						endIdx -= 1
					}
				} else if floatSubSlice > 0.5 && nextIterIdx >= int(intSubSlice) {
					endIdx += 1
				}
			}
		}
		go sortRoutine(sortChannel, userNumbers[startIdx:endIdx])
		startIdx = endIdx
	}

	totalSlices := []int{}
	for i := 0; i < subProcessesAmount; i ++{
		currentSlice := <- sortChannel
		totalSlices = append(totalSlices, currentSlice...)
	}
	fmt.Println("TOTAL sorted list => ", sortNumberSlice(totalSlices))
}

func main(){
	subProcessesAmount := 4
	userNumbers := []int{}
	for {
		var currInput string
		fmt.Println("Type in an integer to be added to the sorted list or 'S' to get the final sorted list:")
		fmt.Scanln(&currInput)
		if currNumber, err := strconv.Atoi(currInput); err == nil {
			userNumbers = append(userNumbers, currNumber)
		} else {
			if strings.ToUpper(currInput) == "S" {
				handleSort(userNumbers, subProcessesAmount)
				break
			} else {
				fmt.Println("The only action in this program is triggered by char 'S'. Please try again.")
			}
		}
	}
}
