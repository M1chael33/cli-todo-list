package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	task      string
	isChecked bool
}

var tasks []data

func addTask(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: add <task>")
		return
	}

	task := strings.Join(parts[1:], " ")
	tasks = append(tasks, data{task: task, isChecked: false})
	fmt.Println("Added:", task)
}

func listTasks() {
	maxLen := 0
	for _, t := range tasks {
		if len(t.task) > maxLen {
			maxLen = len(t.task)
		}
	}

	for i, t := range tasks {
		checkbox := "[ ]"
		if t.isChecked {
			checkbox = "[x]"
		}
		fmt.Printf("%d. %-*s %s\n", i+1, maxLen+2, t.task, checkbox)
	}
}

//  gives a x to signify completion of a task
func checkTask(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: check <task number>")
		return
	}

	num, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error: Please provide a valid number")
		return
	}

	index := num - 1

	if index < 0 || index >= len(tasks) {
		fmt.Println("Error: Task number does not exist")
		return
	}

	tasks[index].isChecked = true
	fmt.Println("Marked done:", tasks[index].task)
}

func delTask(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: remove <task number>")
		return
	}

	num, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error: Please provide a valid number")
		return
	}

	index := num - 1

	if index < 0 || index >= len(tasks) {
		fmt.Println("Error: Task number does not exist")
		return
	}

	removed := tasks[index].task
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Removed:", removed)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "add":
			addTask(parts)
		case "remove":
			delTask(parts)
		case "list":
			listTasks()
		case "check":
			checkTask(parts)
		case "exit", "quit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}