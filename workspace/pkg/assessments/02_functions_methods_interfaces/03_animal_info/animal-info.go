package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func retriveAnimalFromMap(animalName string) Animal{
	animalMap := map[string]Animal {
		"cow": {food: "grass", locomotion: "walk", noise: "moo"},
		"bird": {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}
	return animalMap[animalName]
}

func retrieveAnimalActionFromMap(animal Animal, action string) string{
    actionsMap := map[string]string{
        "eat": animal.Eat(),
        "move": animal.Move(),
        "speak": animal.Speak(),
	}
	return actionsMap[action]
}

func main(){
	postFixErr := "Please try again."
	for {
		var animalName, action string
		fmt.Print("> ")
		fmt.Scanln(&animalName, &action)
		if animalName != "" && action != "" {
			animal := retriveAnimalFromMap(animalName)
			fmt.Println(retrieveAnimalActionFromMap(animal, action))
		} else {
			fmt.Println("Missing target in typed in command: expected <animal> <action>", postFixErr)
		}
	}
}
