package main

import (
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type Todos []Todo

func (todos *Todos) add(message string) {
	todo := Todo{
		Task: message,
		Done: false,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) remove(id int) {

	t := *todos

	*todos = append(t[:id], t[id+1:]...)
}

func (todos *Todos) markDone(id int) {
	t := *todos

	isCompleted := t[id].Done

	if !isCompleted {
		t[id].Done = true

	}

}

func (todos *Todos) edit(id int, task string) {
	t := *todos

	t[id].Task = task

}

func (todos *Todos) clear() {
	*todos = nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Completed")

	for index, t := range *todos {
		completed := "❌"

		if t.Done {
			completed = "✅"
		}

		table.AddRow(strconv.Itoa(index), t.Task, completed)

	}

	table.Render()

}
