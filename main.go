package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type data struct {
	task string
}

var tasks []data

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input) // splits on whitespace
		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <task>")
				continue
			}
			task := strings.Join(parts[1:], " ")
			tasks = append(tasks, data{task: task})
			fmt.Println("Added:", task)

		case "list":
			for i, t := range tasks {
				fmt.Printf("%d: %s\n", i, t.task)
			}

		case "exit", "quit":
			fmt.Println("Bye!")
			return // exits main, ending the program

		default:
			fmt.Println("Unknown command:", command)
		}
	}
}