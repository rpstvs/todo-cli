package main

import (
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

	words := cleanInput(os.Args[1:])

	if len(words) == 0 {
		log.Println("No command, exiting app...")
		return

	}

	command, exists := getCommand()[words[0]]
	task := getMessage(words[1:])

	if exists {
		command.callback(todos, task)
	} else {
		log.Println("unknown command")
	}
}

func cleanInput(text []string) []string {
	var words []string
	for _, field := range text {
		words = append(words, strings.ToLower(field))
	}
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
		"clearCompleted": {
			name:     "clearCompleted",
			desc:     "clears the list of completed tasks",
			callback: commandClearCompleted,
		},
		"print": {
			name:     "print",
			desc:     "prints the list of tasks",
			callback: commandPrint,
		},
		"help": {
			name:     "help",
			desc:     "Lists the commands available",
			callback: commandHelp,
		},
	}
}
func commandPrint(todos *Todos, args ...string) {
	todos.print()
}

func commandHelp(todos *Todos, args ...string) {
	for _, cmd := range getCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.desc)
	}
}

func commandAdd(todos *Todos, args ...string) {
	task := getMessage(args)
	todos.add(task)
	todos.print()
}

func commandRemove(todos *Todos, args ...string) {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatal("couldnt convert to id")
	}
	todos.remove(id)
	todos.print()
}

func commandEdit(todos *Todos, args ...string) {
	words := strings.Fields(args[0])
	id, err := strconv.Atoi(words[0])

	if err != nil {
		log.Fatal("couldnt convert to id")
	}
	task := getMessage(words[1:])
	todos.edit(id, task)
	todos.print()
}

func commandDone(todos *Todos, args ...string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("couldnt convert to id")
	}
	todos.markDone(id)
	todos.print()
}

func commandClear(todos *Todos, args ...string) {
	todos.clear()
}

func commandClearCompleted(todos *Todos, args ...string) {
	for i, todo := range *todos {
		if todo.Done {
			todos.remove(i)
		}
	}
}

func getMessage(msg []string) string {
	task := strings.Join(msg, " ")
	return task
}
