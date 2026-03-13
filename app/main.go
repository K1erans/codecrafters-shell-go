package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		Check_Command(commands, args[1], args[2:])
		return
	case "exit":
		os.Exit(0)
		return
	default:
		Check_Command(commands, args[0], args[1:])
		return
	}
}
func Check_Command(commands []string, args string, args2 []string) {
	if path, err := exec.LookPath(args); err == nil {
		fmt.Println(args + " is " + path)
		command := exec.Command(path, args2...)
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println(string(output))
		}
		return
	} else if slices.Contains(commands, args) {
		fmt.Println(args + " is a shell builtin")
		return
	} else {
		fmt.Println(args + ": not found")
		return
	}
}
