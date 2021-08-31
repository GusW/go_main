package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ftos(floatNum float64) (string, bool) {
	convertedFloat := strconv.FormatFloat(floatNum, 'f', 6, 64)
	if convertedFloat == "" {
		err := fmt.Sprintf("Could not parse float %f to string", floatNum)
		return err, false
	}
	return convertedFloat, true
}

func truncateNumber(floatNum float64) (string, bool) {
	ftosResponse, isSucceeded := ftos(floatNum)
	if isSucceeded {
		targetIdx := strings.Index(ftosResponse, ".")
		switch {
		case targetIdx != -1:
			return ftosResponse[:targetIdx], true
		default:
			warning := fmt.Sprintf("Number %f is not a float. Truncated => %f", floatNum, floatNum)
			return warning, true
		}
	}
	return ftosResponse, false
}

func main() {
	var myFloatNum float64
	fmt.Printf("Enter a decimal whose floating separator is a period '.':\n")
	_, err := fmt.Scan(&myFloatNum)
	if err != nil {
		errMsg := fmt.Sprintf("Error while capturing float: %f => no float found\n", myFloatNum)
		fmt.Print(errMsg)
	} else {
		truncateNumberRes, isSucceeded := truncateNumber(myFloatNum)
		if isSucceeded {
			fmt.Printf("Your truncated number is: %s\n", truncateNumberRes)
		} else {
			fmt.Printf("An error occurred while truncating %f : %s\n", myFloatNum, truncateNumberRes)
		}
	}
}
