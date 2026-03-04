package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}

		Command(strings.Fields(input))
	}
}

func Command(command []string) {
	switch command[0] {
	case "echo":
		var output string
		for _, arg := range command[1:] {
			output += arg + " "
		}
		fmt.Println(output)
		return
	default:
		fmt.Println(command[0] + ": command not found")
	}
}
