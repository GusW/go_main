package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food, locomotion, noise string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food, locomotion, noise string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func retriveAnimalFromMap(animalType string) Animal {
	expectedAnimals := []string{"cow", "bird", "snake"}
	newAnimalType := strings.ToLower(animalType)
	switch newAnimalType {
	case expectedAnimals[0]:
		return Cow{food: "grass", locomotion: "walk", noise: "moo"}
	case expectedAnimals[1]:
		return Bird{food: "worms", locomotion: "fly", noise: "peep"}
	case expectedAnimals[2]:
		return Snake{food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		fmt.Println("Animal type", animalType, "unknown. Expected animal types are:", expectedAnimals)
		return nil
	}
}

func handleAnimalInfo(animal Animal, animalInfo string) {
	expectedAnimalInfo := []string{"eat", "move", "speak"}
	newAnimalInfo := strings.ToLower(animalInfo)
	switch newAnimalInfo {
	case expectedAnimalInfo[0]:
		animal.Eat()
	case expectedAnimalInfo[1]:
		animal.Move()
	case expectedAnimalInfo[2]:
		animal.Speak()
	default:
		fmt.Println("Animal info", animalInfo, "unknown. Expected animal infos are:", expectedAnimalInfo)
	}
}

type AnimalAction func(animals map[string]Animal, args ...string)

func findAnimalByName(animals map[string]Animal, animalName string) Animal {
	if animal, ok := animals[animalName]; ok {
		return animal
	} else {
		return nil
	}
}

func createAnimal(animals map[string]Animal, args ...string) {
	animalName, animalType := args[0], args[1]
	preexistingAnimal := findAnimalByName(animals, animalName)
	if preexistingAnimal != nil {
		fmt.Println("There is an animal named", animalName, "already, punk. Please choose another name and try again.")
	} else {
		newAnimal := retriveAnimalFromMap(animalType)
		if newAnimal != nil {
			animals[animalName] = newAnimal
			fmt.Println("Created it!")
		}
	}
}

func queryAnimal(animals map[string]Animal, args ...string) {
	animalName := args[0]
	animal := findAnimalByName(animals, animalName)
	if animal != nil {
		animalInfo := args[1]
		handleAnimalInfo(animal, animalInfo)
	} else {
		fmt.Println("Animal name", animalName, "unknown. Please try again.")
	}
}

func handleAction(animals map[string]Animal, action string, args ...string) {
	expectedActions := []string{"newanimal", "query"}
	newAction := strings.ToLower(action)
	actionMap := map[string]AnimalAction{
		expectedActions[0]: createAnimal,
		expectedActions[1]: queryAnimal,
	}
	if actionTarget, ok := actionMap[newAction]; ok {
		actionTarget(animals, args...)
	} else {
		fmt.Println("Action", action, "unknown. Expected actions are:", expectedActions)
	}
}

func main() {
	var action, name, type_ string
	animals := make(map[string]Animal)
	for {
		fmt.Print("> ")
		fmt.Scanln(&action, &name, &type_)
		if action != "" && name != "" && type_ != "" {
			handleAction(animals, action, name, type_)
		} else {
			fmt.Println("Expected 3 inputs <action> <name> <type>. Please try again.")
		}
		action = ""
		name = ""
		type_ = ""
	}
}
