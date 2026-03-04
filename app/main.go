package main

import (
	"fmt"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input, ": command not found")
}
