package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	name     string
	desc     string
	callback func(todos *Todos, args ...string)
}

func processArgs(todos *Todos) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	words := cleanInput(scanner.Text())

	if len(words) == 0 {
		log.Println("No command, exiting app...")
		return

	}

	command, exists := getCommand()[words[0]]

	if exists {
		command.callback(todos, words[1:]...)
	} else {
		log.Println("unknown command")
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommand() map[string]Command {
	return map[string]Command{
		"add": {
			name:     "add",
			desc:     "adds a task to the list",
			callback: commandAdd,
		},
		"remove": {
			name:     "remove",
			desc:     "removes item from the list",
			callback: commandRemove,
		},
		"done": {
			name:     "done",
			desc:     "mark task as done",
			callback: commandDone,
		},
		"edit": {
			name:     "edit",
			desc:     "edits a specific task",
			callback: commandEdit,
		},
		"clear": {
			name:     "clear",
			desc:     "clears the list",
			callback: commandClear,
		},
	}
}

func commandAdd(todos *Todos, args ...string) {
	todos.add(args[0])
	todos.print()
}

func commandRemove(todos *Todos, args ...string) {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatal("couldnt convert to id")
	}
	todos.remove(id)
}

func commandEdit(todos *Todos, args ...string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("couldnt convert to id")
	}

	todos.edit(id, args[1])
}

func commandDone(todos *Todos, args ...string) {
	id, err := strconv.Atoi(args[0])
	fmt.Println("estou a entrar aqui")
	if err != nil {
		log.Fatal("couldnt convert to id")
	}
	todos.markDone(id)
	todos.print()
}

func commandClear(todos *Todos, args ...string) {
	todos.clear()
}
