package main

import (
	"fmt"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	var input string
	for {
		fmt.Print("$ ")
		fmt.Scanln(&input)
		switch input {
		case "exit":
			return
		default:
			fmt.Println(input + ": command not found")
		}
	}
}
