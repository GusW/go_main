package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// file should be alike
// first_name_1 last_name_1
// first_name_2 last_name_2
// ...
// first_name_n last_name_n

type Name struct {
	fname, lname string
}

func _captureFilePath(filePath *string) bool {
	fmt.Print("ATTENTION: behavior must change in different terminals\n")
	fmt.Print("ATTENTION: tested on LINUX BASH only\n")
	fmt.Print("Type in an /absolute/path/to/file.txt:\n")
	_, errInput := fmt.Scan(filePath)
	if errInput != nil {
		fmt.Printf("Could not capture file path: err => %s\n", errInput)
		return true
	}
	return false
}

func _handleFileContents(filePath string) string {
	content, errRead := ioutil.ReadFile(filePath)
	if errRead != nil {
		fmt.Printf("Could not read file: err => %s\n", errRead)
		return ""
	}
	return string(content)
}

func _generateNameObjs(contents string) []Name {
	var names = make([]Name, 0)
	// code smelly, but I dont want to use bufio.ReadLine ;)
	nameArray := strings.Split(contents, "\n")
	for i := 0; i < len(nameArray); i++ {
		name := strings.Split(nameArray[i], " ")
		if len(name) == 2 {
			currentName := Name{fname: name[0], lname: name[1]}
			names = append(names, currentName)
		}
	}

	return names
}

func main() {

	var filePath string
	hasInputErr := _captureFilePath(&filePath)
	if hasInputErr {
		return
	}
	fileContents := _handleFileContents(filePath)
	if fileContents == "" {
		return
	}
	names := _generateNameObjs(fileContents)
	if len(names) > 0 {
		fmt.Print("Capture name structs from file are:\n", names, "\n")
	} else {
		fmt.Print("No name struct was capture from the file in", filePath, "\n")
	}
}
