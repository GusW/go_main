package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func GenAnimalFn(food, locomotion, noise string) Animal {
	return Animal{
		food:       food,
		locomotion: locomotion,
		noise:      noise,
	}
}

func main() {
	animals := make(map[string]Animal)
	animals["cow"] = GenAnimalFn("grass", "walk", "moo")
	animals["bird"] = GenAnimalFn("worms", "fly", "peep")
	animals["snake"] = GenAnimalFn("mice", "slither", "hsss")

	reader := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Print("> ")
		s, _ = reader.ReadString('\n')
		s = strings.Replace(s, "\r\n", "", -1)
		s = strings.Replace(s, "\n", "", -1)
		cmd := strings.Split(s, " ")

		subj := strings.ToLower(cmd[0])
		verb := strings.ToLower(cmd[1])

		a, ok := animals[subj]

		if !ok {
			fmt.Println("Undefined animal")
		}

		switch verb {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			fmt.Println("Undefined verb")
		}
	}

}
