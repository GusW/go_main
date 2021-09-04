package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var s1, s2 bool
	var str string
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Received Error", err)
	}
	str = strings.ToLower(line)
	r := regexp.MustCompile(`(?m)^i[a-z0-9 ]+n$`)
	s1 = r.MatchString(str)
	s2 = strings.ContainsAny(str, "a")
	if s1 && s2 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	fmt.Printf(" Dump input : %s  s1 : %t s2: %t ", str, s1, s2)

}
