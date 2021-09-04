package main

import (
	"encoding/json"
	"fmt"
)


func _handleUserInput(name *string, addr *string) (string){
	fmt.Print("Enter user name:\n")
	_, errName := fmt.Scan(name)
	if errName != nil {
		return errName.Error()
	}
	fmt.Print("Enter user address:\n")
	_, errAddr := fmt.Scan(addr)
	if errAddr != nil {
		return errName.Error()
	}
	return ""
}


func main(){
	var name, addr string
	errInput := _handleUserInput(&name, &addr)
	if errInput != "" {
		fmt.Printf("Could not capture user input: error => %s", errInput)
	} else {
		userMap := map[string]string{"name": name, "address": addr}
		userMapByteArr, errJson := json.Marshal(userMap)
		if errJson != nil {
			fmt.Printf("Could not marshal user map: error => %s", errJson)
		} else {
			userJson := string(userMapByteArr)
			fmt.Printf("JSON representation for user name %s address %s is:\n", name, addr)
			fmt.Print(userJson, "\n")
		}
	}
}
