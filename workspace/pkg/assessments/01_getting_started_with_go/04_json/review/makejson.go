package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	inReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	inName, err := inReader.ReadString('\n')
	if err != nil {
		panic("Can't read line")
	}

	fmt.Print("Enter address: ")
	inAddress, err := inReader.ReadString('\n')
	if err != nil {
		panic("Can't read line")
	}

	addressMap := map[string]string{
		"name":    strings.TrimSpace(inName),
		"address": strings.TrimSpace(inAddress),
	}
	jsonbytes, _ := json.Marshal(addressMap)
	fmt.Printf("Result: %s", jsonbytes)
}
