package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	reader := bufio.NewReader(os.Stdin)
	commands := []string{"echo", "type", "exit"}
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		Command(strings.Fields(input), commands)
	}
}

func Command(args []string, commands []string) {
	switch args[0] {
	case "echo":
		var output string
		for _, arg := range args[1:] {
			output += arg + " "
		}
		fmt.Println(output)
		return
	case "type":
		if slices.Contains(commands, args[1]) {
			fmt.Println(args[1] + " is a shell builtin")
		} else {
			fmt.Println(args[1] + ": not found")
		}
		return
	case "exit":
		os.Exit(1)
		return
	default:
		fmt.Println(args[0] + ": command not found")
	}
}
