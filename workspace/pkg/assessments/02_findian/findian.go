package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func _castFountRegexpToMessage(matched bool) string {
	if matched {
		return "Found!\n"
	}
	return "Not Found!\n"
}

func _findRegex(sourceString string) (string, bool) {
	matched, err := regexp.MatchString(`^[iI](.*)[aA](.*)[nN]$`, sourceString)
	if err != nil {
		return _castFountRegexpToMessage(matched), false
	}
	return _castFountRegexpToMessage(matched), true
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Type in anything:\n")

	userInput, err := consoleReader.ReadString('\n')
	if err != nil {
		errMsg := fmt.Sprintf("Could not capture user input: %s\n", err)
		fmt.Print(errMsg)
	} else {
		trimmedUserInput := strings.Trim(userInput, " \n")
		res, isSucceeded := _findRegex(trimmedUserInput)
		if isSucceeded {
			fmt.Print(res)
		} else {
			warningMsg := fmt.Sprintf("Could not handle regex search on string %s: %s\n", userInput, res)
			fmt.Print(warningMsg)
		}
	}
}
