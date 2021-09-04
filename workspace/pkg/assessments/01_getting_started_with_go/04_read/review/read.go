package main

import (
	//"encoding/json"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	fmt.Println("Enter the path to the file to read:")
	var fileName string
	fmt.Scanln(&fileName)

	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	names := make([]Person, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		personNames := strings.Split(line, " ")
		if len(personNames) < 2 {
			fmt.Println("Error, unexpected number of words in line")
			os.Exit(1)
		}
		person := Person{
			fname: personNames[0],
			lname: personNames[1],
		}
		names = append(names, person)
	}

	for _, name := range names {
		fmt.Printf("%s, %s \n", name.fname, name.lname)
	}
}
