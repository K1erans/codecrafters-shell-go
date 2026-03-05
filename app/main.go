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
		if slices.Contains(commands, args[1]) {
			fmt.Println(args[1] + " is a shell builtin")
		} else if path, err := exec.LookPath(args[1]); err == nil {
			fmt.Println(args[1] + " is " + path)
		}
		return
	case "exit":
		os.Exit(0)
		return
	default:
		fmt.Println(args[0] + ": command not found")
	}
}

// func Path_Traversal(path string, arg string) string {
// 	var hard_to_name_variable string = ""
// 	parts := strings.Split(path, "/")
// 	for _, part := range parts {
// 		var part_path string = path + "/" + part
// 		file, err := os.Open(part_path)
// 		if err == nil {
// 			hard_to_name_variable = arg + " is " + part
// 			return hard_to_name_variable
// 		}
// 		file.Close()
// 	}
//   hard_to_name_variable = arg + ": command not found"
//   return hard_to_name_variable
// }
